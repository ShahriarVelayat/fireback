import { CommonSingleManager } from "@/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/fireback/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/fireback/hooks/useCommonEntityManager";
import { useT } from "@/fireback/hooks/useT";
import { useGetShoppingCartByUniqueId } from "src/sdk/fireback/modules/shop/useGetShoppingCartByUniqueId";
import { ShoppingCartEntity } from "src/sdk/fireback/modules/shop/ShoppingCartEntity";
export const ShoppingCartSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});
  const getSingleHook = useGetShoppingCartByUniqueId({ query: { uniqueId } });
  var d: ShoppingCartEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);
  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(ShoppingCartEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView entity={d} fields={[]} />
      </CommonSingleManager>
    </>
  );
};
