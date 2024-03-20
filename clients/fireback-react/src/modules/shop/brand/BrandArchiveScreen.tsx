import { useT } from "@/hooks/useT";
import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { BrandList } from "./BrandList";
import { BrandEntity } from "src/sdk/fireback/modules/shop/BrandEntity";
export const BrandArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.brands.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(BrandEntity.Navigation.create(locale));
      }}
    >
      <BrandList />
    </CommonArchiveManager>
  );
};