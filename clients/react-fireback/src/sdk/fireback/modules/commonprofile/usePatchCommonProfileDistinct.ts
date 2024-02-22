// @ts-nocheck

import { FormikHelpers } from "formik";
import React, { useCallback, useContext } from "react";
import {
  useMutation,
  useQuery,
  useQueryClient,
  QueryClient,
} from "react-query";
import { CommonProfileActions } from "./common-profile-actions";
import * as commonprofile from "./index";
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

export function usePatchCommonProfileDistinct({
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
    ? CommonProfileActions.fnExec(execFnOverride(options))
    : execFn
    ? CommonProfileActions.fnExec(execFn(options))
    : CommonProfileActions.fn(options);
  const Q = () => fnx;

  const fn = (entity: any) => Q().patchCommonProfileDistinct(entity);

  const mutation = useMutation<
    IResponse<commonprofile.CommonProfileEntity>,
    IResponse<commonprofile.CommonProfileEntity>,
    Partial<commonprofile.CommonProfileEntity>
  >(fn);

  // Only entities are having a store in front-end

  const fnUpdater = (
    data: IResponseList<commonprofile.CommonProfileEntity> | undefined,
    item: IResponse<commonprofile.CommonProfileEntity>
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
    //       CommonProfileActions.isCommonProfileEntityEqual(t, item.data)
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
    values: Partial<commonprofile.CommonProfileEntity>,
    formikProps?: FormikHelpers<Partial<commonprofile.CommonProfileEntity>>
  ): Promise<IResponse<commonprofile.CommonProfileEntity>> => {
    return new Promise((resolve, reject) => {
      mutation.mutate(values, {
        onSuccess(response: IResponse<commonprofile.CommonProfileEntity>) {
          queryClient.setQueriesData(
            "*commonprofile.CommonProfileEntity",
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
