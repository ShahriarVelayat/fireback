import { CommonSingleManager } from "@/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import { useT } from "@/fireback/hooks/useT";
import { useGetBrandByUniqueId } from "src/sdk/fireback/modules/shop/useGetBrandByUniqueId";
import { BrandEntity } from "src/sdk/fireback/modules/shop/BrandEntity";
export const BrandSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetBrandByUniqueId({ query: { uniqueId } });
  var d: BrandEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(BrandEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              elem: d?.name,
              label: t.brands.name,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
