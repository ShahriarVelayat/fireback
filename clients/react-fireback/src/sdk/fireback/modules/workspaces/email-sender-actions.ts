// @ts-nocheck
/**
 *  This is an auto generate file from fireback project.
 *  You can use this in order to communicate in backend, it gives you available actions,
 *  and their types
 *  Module: workspaces
 */

import * as workspaces from "./index";
import {
  execApiFn,
  RemoteRequestOption,
  IDeleteResponse,
  IResponse,
  core,
  ExecApi,
  IResponseList,
} from "../../core/http-tools";

export class EmailSenderActions {
  private _itemsPerPage?: number = undefined;
  private _startIndex?: number = undefined;
  private _sort?: number = undefined;
  private _query?: string = undefined;
  private _jsonQuery?: any = undefined;
  private _withPreloads?: string = undefined;
  private _uniqueId?: string = undefined;
  private _deep?: boolean = undefined;

  constructor(private apiFn: ExecApi) {
    this.apiFn = apiFn;
  }

  static isEmailSenderEntityEqual(
    a: workspaces.EmailSenderEntity,
    b: workspaces.EmailSenderEntity
  ): boolean {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId === b.uniqueId;
  }

  static getEmailSenderEntityPrimaryKey(
    a: workspaces.EmailSenderEntity
  ): string {
    // Change if the primary key is different, or is combined with few fields
    return a.uniqueId;
  }

  query(complexSqlAlike: string): EmailSenderActions {
    this._query = complexSqlAlike;
    return this;
  }

  static fn(options: RemoteRequestOption): EmailSenderActions {
    return new EmailSenderActions(execApiFn(options));
  }

  static fnExec(fn: ExecApi): EmailSenderActions {
    return new EmailSenderActions(fn);
  }

  uniqueId(id: string): EmailSenderActions {
    this._uniqueId = id;
    return this;
  }

  deep(deep = true): EmailSenderActions {
    this._deep = deep;
    return this;
  }

  withPreloads(withPreloads: string): EmailSenderActions {
    this._withPreloads = withPreloads;
    return this;
  }

  jsonQuery(q: any): EmailSenderActions {
    this._jsonQuery = q;
    return this;
  }

  sort(sortFields: string | string[]): EmailSenderActions {
    this._sort = sortFields;
    return this;
  }

  startIndex(offset: number): EmailSenderActions {
    this._startIndex = offset;
    return this;
  }

  itemsPerPage(limit: number): EmailSenderActions {
    this._itemsPerPage = limit;
    return this;
  }

  get paramsAsString(): string {
    const q: any = {
      startIndex: this._startIndex,
      itemsPerPage: this._itemsPerPage,
      query: this._query,
      deep: this._deep,
      jsonQuery: JSON.stringify(this._jsonQuery),
      withPreloads: this._withPreloads,
      uniqueId: this._uniqueId,
      sort: this._sort,
    };

    const searchParams = new URLSearchParams();
    Object.keys(q).forEach((key) => {
      if (q[key]) {
        searchParams.append(key, q[key]);
      }
    });

    return searchParams.toString();
  }

  async getEmailSenders(): Promise<
    IResponseList<workspaces.EmailSenderEntity>
  > {
    return this.apiFn(
      "GET",
      `emailSenders?action=EmailSenderActionQuery&${this.paramsAsString}`
    );
  }

  async getEmailSendersExport(): Promise<
    IResponseList<workspaces.EmailSenderEntity>
  > {
    return this.apiFn(
      "GET",
      `emailSenders/export?action=EmailSenderActionExport&${this.paramsAsString}`
    );
  }

  async getEmailSenderByUniqueId(
    uniqueId: string
  ): Promise<IResponse<workspaces.EmailSenderEntity>> {
    return this.apiFn(
      "GET",
      `emailSender/:uniqueId?action=EmailSenderActionGetOne&${this.paramsAsString}`.replace(
        ":uniqueId",
        uniqueId
      )
    );
  }

  async postEmailSender(
    entity: workspaces.EmailSenderEntity
  ): Promise<IResponse<workspaces.EmailSenderEntity>> {
    return this.apiFn(
      "POST",
      `emailSender?action=EmailSenderActionCreate&${this.paramsAsString}`,
      entity
    );
  }

  async patchEmailSender(
    entity: workspaces.EmailSenderEntity
  ): Promise<IResponse<workspaces.EmailSenderEntity>> {
    return this.apiFn(
      "PATCH",
      `emailSender?action=EmailSenderActionUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async patchEmailSenders(
    entity: core.BulkRecordRequest<workspaces.EmailSenderEntity>
  ): Promise<IResponse<core.BulkRecordRequest[workspaces.EmailSenderEntity]>> {
    return this.apiFn(
      "PATCH",
      `emailSenders?action=EmailSenderActionBulkUpdate&${this.paramsAsString}`,
      entity
    );
  }

  async deleteEmailSender(
    entity: core.DeleteRequest
  ): Promise<IDeleteResponse> {
    return this.apiFn(
      "DELETE",
      `emailSender?action=EmailSenderActionRemove&${this.paramsAsString}`,
      entity
    );
  }
}
