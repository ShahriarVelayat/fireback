import { useT } from "@/hooks/useT";
import { CommonArchiveManager } from "@/components/entity-manager/CommonArchiveManager";
import { CategoryList } from "./CategoryList";
import { CategoryEntity } from "src/sdk/fireback/modules/shop/CategoryEntity";
export const CategoryArchiveScreen = () => {
  const t = useT();
  return (
    <CommonArchiveManager
      pageTitle={t.categories.archiveTitle}
      newEntityHandler={({ locale, router }) => {
        router.push(CategoryEntity.Navigation.create(locale));
      }}
    >
      <CategoryList />
    </CommonArchiveManager>
  );
};