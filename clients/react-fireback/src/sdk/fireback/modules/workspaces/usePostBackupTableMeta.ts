// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { BackupTableMetaActions } from "./backup-table-meta-actions";
import * as workspaces from "./index";
import {
  execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  core,
  IResponse,
  ExecApi,
  mutationErrorsToFormik,
  IResponseList,
} from "../../core/http-tools";
import { RemoteQueryContext } from "../../core/react-tools";

export function usePostBackupTableMeta({
  queryClient,
  query,
  execFnOverride,
}: {
  queryClient: QueryClient;
  query?: any;
  execFnOverride?: any;
}) {
  query = query || {};

  const { options, execFn } = useContext(RemoteQueryContext);

  const fnx = execFnOverride
    ? BackupTableMetaActions.fnExec(execFnOverride(options))
    : execFn
    ? BackupTableMetaActions.fnExec(execFn(options))
    : BackupTableMetaActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().postBackupTableMeta(entity);

  const mutation = useMutation<
    IResponse<workspaces.BackupTableMetaEntity>,
    IResponse<workspaces.BackupTableMetaEntity>,
    Partial<workspaces.BackupTableMetaEntity>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<workspaces.BackupTableMetaEntity> | undefined,
    item: IResponse<workspaces.BackupTableMetaEntity>
  ) => {
    if (!data) {
      return {
        data: { items: [] },
      };
    }

    // To me it seems this is not a good or any correct strategy to update the store.
    // When we are posting, we want to add it there, that's it. Not updating it.
    // We have patch, but also posting with ID is possible.

    // if (data?.data?.items && item.data) {
    //   data.data.items = data.data.items.map((t) => {
    //     if (
    //       item.data !== undefined &&
    //       BackupTableMetaActions.isBackupTableMetaEntityEqual(t, item.data)
    //     ) {
    //       return item.data;
    //     }

    //     return t;
    //   });
    // } else if (data?.data && item.data) {
    //   data.data.items = [item.data, ...(data?.data?.items || [])];
    // }

    data.data.items = [item.data, ...(data?.data?.items || [])];

    return data;
  };

  const submit = (
    values: Partial<workspaces.BackupTableMetaEntity>,
    formikProps?: FormikHelpers<Partial<workspaces.BackupTableMetaEntity>>
  ): Promise<IResponse<workspaces.BackupTableMetaEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<workspaces.BackupTableMetaEntity>) {
          queryClient.setQueryData<
            IResponseList<workspaces.BackupTableMetaEntity>
          >("*workspaces.BackupTableMetaEntity", (data) =>
            fnUpdater(data, response)
          );

          resolve(response);
        },

        onError(error: any) {
          formikProps?.setErrors(mutationErrorsToFormik(error));

          reject(error);
        },
      });
    });
  };

  return { mutation, submit, fnUpdater };
}
