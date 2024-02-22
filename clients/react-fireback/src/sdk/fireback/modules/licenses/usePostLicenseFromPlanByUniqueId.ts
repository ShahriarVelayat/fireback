// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { LicenseActions } from "./license-actions";
import * as licenses from "./index";
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

export function usePostLicenseFromPlanByUniqueId({
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
    ? LicenseActions.fnExec(execFnOverride(options))
    : execFn
    ? LicenseActions.fnExec(execFn(options))
    : LicenseActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) =>
    Q().postLicenseFromPlanByUniqueId(
      query.uniqueId,

      entity
    );

  const mutation = useMutation<
    IResponse<licenses.LicenseEntity>,
    IResponse<licenses.LicenseEntity>,
    Partial<licenses.LicenseFromPlanIdDto>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<licenses.LicenseEntity> | undefined,
    item: IResponse<licenses.LicenseEntity>
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
    //       LicenseActions.isLicenseEntityEqual(t, item.data)
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
    values: Partial<licenses.LicenseFromPlanIdDto>,
    formikProps?: FormikHelpers<Partial<licenses.LicenseEntity>>
  ): Promise<IResponse<licenses.LicenseEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<licenses.LicenseEntity>) {
          queryClient.setQueryData<IResponseList<licenses.LicenseEntity>>(
            "*licenses.LicenseEntity",
            (data) => fnUpdater(data, response)
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
