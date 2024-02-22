/**
 * Fields for entity. Use this when creating forms in React/Angular,
 * instead of giving string to each one, use it from here, so in case of
 * updating any fields you don't loose it.
 */

export const WorkspaceTypeEntityFields = {
  visibility: "visibility",
  workspaceId: "workspaceId",
  linkerId: "linkerId",
  parentId: "parentId",
  parent: "parent",
  parent$: "parent",
  uniqueId: "uniqueId",
  userId: "userId",
  translations: "translations",
  title: "title",
  description: "description",
  slug: "slug",
  roleId: "roleId",
  role: {
    visibility: "role.visibility",
    workspaceId: "role.workspaceId",
    linkerId: "role.linkerId",
    parentId: "role.parentId",
    parent: "role.parent",
    parent$: "parent",
    uniqueId: "role.uniqueId",
    userId: "role.userId",
    name: "role.name",
    capabilitiesListId: "role.capabilitiesListId",
    capabilities: {
      visibility: "role.capabilities.visibility",
      workspaceId: "role.capabilities.workspaceId",
      linkerId: "role.capabilities.linkerId",
      parentId: "role.capabilities.parentId",
      parent: "role.capabilities.parent",
      parent$: "parent",
      uniqueId: "role.capabilities.uniqueId",
      userId: "role.capabilities.userId",
      name: "role.capabilities.name",
      rank: "role.capabilities.rank",
      updated: "role.capabilities.updated",
      created: "role.capabilities.created",
      createdFormatted: "role.capabilities.createdFormatted",
      updatedFormatted: "role.capabilities.updatedFormatted",
    },
    capabilities$: "capabilities",
    rank: "role.rank",
    updated: "role.updated",
    created: "role.created",
    createdFormatted: "role.createdFormatted",
    updatedFormatted: "role.updatedFormatted",
  },
  role$: "role",
  rank: "rank",
  updated: "updated",
  created: "created",
  createdFormatted: "createdFormatted",
  updatedFormatted: "updatedFormatted",
};
