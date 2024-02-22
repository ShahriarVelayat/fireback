// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { PassportMethodActions } from "./passport-method-actions";
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

export function usePostPassportMethod({
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
    ? PassportMethodActions.fnExec(execFnOverride(options))
    : execFn
    ? PassportMethodActions.fnExec(execFn(options))
    : PassportMethodActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().postPassportMethod(entity);

  const mutation = useMutation<
    IResponse<workspaces.PassportMethodEntity>,
    IResponse<workspaces.PassportMethodEntity>,
    Partial<workspaces.PassportMethodEntity>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<workspaces.PassportMethodEntity> | undefined,
    item: IResponse<workspaces.PassportMethodEntity>
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
    //       PassportMethodActions.isPassportMethodEntityEqual(t, item.data)
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
    values: Partial<workspaces.PassportMethodEntity>,
    formikProps?: FormikHelpers<Partial<workspaces.PassportMethodEntity>>
  ): Promise<IResponse<workspaces.PassportMethodEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<workspaces.PassportMethodEntity>) {
          queryClient.setQueryData<
            IResponseList<workspaces.PassportMethodEntity>
          >("*workspaces.PassportMethodEntity", (data) =>
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