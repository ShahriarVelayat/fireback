import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { useGetPageCategoryByUniqueId } from "src/sdk/fireback/modules/cms/useGetPageCategoryByUniqueId";
import { PageCategoryEntity } from "src/sdk/fireback/modules/cms/PageCategoryEntity";
export const PageCategorySingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetPageCategoryByUniqueId({ query: { uniqueId } });
  var d: PageCategoryEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(PageCategoryEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={
            [
              {
                elem: d?.name,
                label: t.pagecategories.name,
              },    
            ]
          }
        />
      </CommonSingleManager>
    </>
  );
};