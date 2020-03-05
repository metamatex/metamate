export default sdk

namespace sdk {
    
    export interface AttachmentFilter {
        alternativeIds?: IdListFilter;
        and?: AttachmentFilter[];
        description?: StringFilter;
        hash?: string;
        id?: ServiceIdFilter;
        meta?: TypeMetaFilter;
        not?: AttachmentFilter[];
        or?: AttachmentFilter[];
        set?: boolean;
    }
    
    export interface AttachmentListFilter {
        every?: AttachmentFilter;
        hash?: string;
        none?: AttachmentFilter;
        some?: AttachmentFilter;
    }
    
    export interface Auth {
        clientAccount?: ClientAccount;
        hash?: string;
        token?: Token;
    }
    
    export interface AuthenticateClientAccountEndpoint {
        filter?: AuthenticateClientAccountRequestFilter;
        hash?: string;
    }
    
    export interface AuthenticateClientAccountEndpointFilter {
        and?: AuthenticateClientAccountEndpointFilter[];
        hash?: string;
        not?: AuthenticateClientAccountEndpointFilter[];
        or?: AuthenticateClientAccountEndpointFilter[];
        set?: boolean;
    }
    
    export interface AuthenticateClientAccountInputFilter {
        and?: AuthenticateClientAccountInputFilter[];
        hash?: string;
        id?: IdFilter;
        not?: AuthenticateClientAccountInputFilter[];
        or?: AuthenticateClientAccountInputFilter[];
        password?: StringFilter;
        set?: boolean;
    }
    
    export interface AuthenticateClientAccountRequestFilter {
        and?: AuthenticateClientAccountRequestFilter[];
        hash?: string;
        input?: AuthenticateClientAccountInputFilter;
        meta?: RequestMetaFilter;
        not?: AuthenticateClientAccountRequestFilter[];
        or?: AuthenticateClientAccountRequestFilter[];
        set?: boolean;
    }
    
    export interface BlueWhatever {
        alternativeIds?: Id[];
        boolField?: boolean;
        boolList?: boolean[];
        enumField?: string;
        enumList?: string[];
        float64Field?: number;
        float64List?: number[];
        hash?: string;
        id?: ServiceId;
        int32Field?: number;
        int32List?: number[];
        meta?: TypeMeta;
        relations?: BlueWhateverRelations;
        stringField?: string;
        stringList?: string[];
        unionField?: WhateverUnion;
        unionList?: WhateverUnion[];
    }
    
    export interface BlueWhateverFilter {
        alternativeIds?: IdListFilter;
        and?: BlueWhateverFilter[];
        boolField?: BoolFilter;
        boolList?: BoolListFilter;
        enumField?: EnumFilter;
        enumList?: EnumListFilter;
        float64Field?: Float64Filter;
        float64List?: Float64ListFilter;
        hash?: string;
        id?: ServiceIdFilter;
        int32Field?: Int32Filter;
        int32List?: Int32ListFilter;
        meta?: TypeMetaFilter;
        not?: BlueWhateverFilter[];
        or?: BlueWhateverFilter[];
        set?: boolean;
        stringField?: StringFilter;
        stringList?: StringListFilter;
        unionField?: WhateverUnionFilter;
        unionList?: WhateverUnionListFilter;
    }
    
    export interface BlueWhateverListFilter {
        every?: BlueWhateverFilter;
        hash?: string;
        none?: BlueWhateverFilter;
        some?: BlueWhateverFilter;
    }
    
    export interface BlueWhateverRelations {
        hash?: string;
        knewByWhatevers?: WhateversCollection;
    }
    
    export interface BlueWhateverRelationsSelect {
        hash?: string;
        knewByWhatevers?: WhateversCollectionSelect;
        selectAll?: boolean;
    }
    
    export interface BlueWhateversCollection {
        blueWhatevers?: BlueWhatever[];
        hash?: string;
        meta?: CollectionMeta;
    }
    
    export interface BlueWhateversCollectionSelect {
        blueWhatevers?: BlueWhateverSelect;
        hash?: string;
        meta?: CollectionMetaSelect;
        selectAll?: boolean;
    }
    
    export interface BlueWhateverSelect {
        alternativeIds?: IdSelect;
        boolField?: boolean;
        boolList?: boolean;
        enumField?: boolean;
        enumList?: boolean;
        float64Field?: boolean;
        float64List?: boolean;
        hash?: string;
        id?: ServiceIdSelect;
        int32Field?: boolean;
        int32List?: boolean;
        meta?: TypeMetaSelect;
        relations?: BlueWhateverRelationsSelect;
        selectAll?: boolean;
        stringField?: boolean;
        stringList?: boolean;
        unionField?: WhateverUnionSelect;
        unionList?: WhateverUnionSelect;
    }
    
    export interface BlueWhateverSort {
        boolField?: string;
        float64Field?: string;
        hash?: string;
        id?: ServiceIdSort;
        int32Field?: string;
        meta?: TypeMetaSort;
        stringField?: string;
        unionField?: WhateverUnionSort;
    }
    
    export interface BoolFilter {
        and?: BoolFilter[];
        hash?: string;
        is?: boolean;
        not?: boolean;
        or?: BoolFilter[];
        set?: boolean;
    }
    
    export interface BoolListFilter {
        and?: BoolFilter;
        hash?: string;
        not?: BoolFilter;
        or?: BoolFilter;
    }
    
    export interface ClientAccount {
        alternativeIds?: Id[];
        hash?: string;
        id?: ServiceId;
        meta?: TypeMeta;
        password?: Password;
        relations?: ClientAccountRelations;
    }
    
    export interface ClientAccountFilter {
        alternativeIds?: IdListFilter;
        and?: ClientAccountFilter[];
        hash?: string;
        id?: ServiceIdFilter;
        meta?: TypeMetaFilter;
        not?: ClientAccountFilter[];
        or?: ClientAccountFilter[];
        password?: PasswordFilter;
        set?: boolean;
    }
    
    export interface ClientAccountListFilter {
        every?: ClientAccountFilter;
        hash?: string;
        none?: ClientAccountFilter;
        some?: ClientAccountFilter;
    }
    
    export interface ClientAccountRelations {
        hash?: string;
        ownsServiceAccounts?: ServiceAccountsCollection;
    }
    
    export interface ClientAccountRelationsSelect {
        hash?: string;
        ownsServiceAccounts?: ServiceAccountsCollectionSelect;
        selectAll?: boolean;
    }
    
    export interface ClientAccountsCollection {
        clientAccounts?: ClientAccount[];
        hash?: string;
        meta?: CollectionMeta;
    }
    
    export interface ClientAccountsCollectionSelect {
        clientAccounts?: ClientAccountSelect;
        hash?: string;
        meta?: CollectionMetaSelect;
        selectAll?: boolean;
    }
    
    export interface ClientAccountSelect {
        alternativeIds?: IdSelect;
        hash?: string;
        id?: ServiceIdSelect;
        meta?: TypeMetaSelect;
        password?: PasswordSelect;
        relations?: ClientAccountRelationsSelect;
        selectAll?: boolean;
    }
    
    export interface ClientAccountSort {
        hash?: string;
        id?: ServiceIdSort;
        meta?: TypeMetaSort;
        password?: PasswordSort;
    }
    
    export interface CollectionGetMode {
        hash?: string;
        pages?: ServicePage[];
    }
    
    export interface CollectionGetModeFilter {
        and?: CollectionGetModeFilter[];
        hash?: string;
        not?: CollectionGetModeFilter[];
        or?: CollectionGetModeFilter[];
        pages?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface CollectionMeta {
        count?: number;
        errors?: Error[];
        hash?: string;
        pagination?: Pagination;
    }
    
    export interface CollectionMetaFilter {
        and?: CollectionMetaFilter[];
        count?: Int32Filter;
        errors?: ErrorListFilter;
        hash?: string;
        not?: CollectionMetaFilter[];
        or?: CollectionMetaFilter[];
        pagination?: PaginationFilter;
        set?: boolean;
    }
    
    export interface CollectionMetaSelect {
        count?: boolean;
        errors?: ErrorSelect;
        hash?: string;
        pagination?: PaginationSelect;
        selectAll?: boolean;
    }
    
    export interface CollectionPostMode {
        hash?: string;
    }
    
    export interface ContextPipeModeFilter {
        and?: ContextPipeModeFilter[];
        hash?: string;
        method?: EnumFilter;
        not?: ContextPipeModeFilter[];
        or?: ContextPipeModeFilter[];
        requester?: EnumFilter;
        set?: boolean;
        stage?: EnumFilter;
    }
    
    export interface CursorPage {
        hash?: string;
        value?: string;
    }
    
    export interface CursorPageFilter {
        and?: CursorPageFilter[];
        hash?: string;
        not?: CursorPageFilter[];
        or?: CursorPageFilter[];
        set?: boolean;
        value?: StringFilter;
    }
    
    export interface CursorPageSelect {
        hash?: string;
        selectAll?: boolean;
        value?: boolean;
    }
    
    export interface DeleteAttachmentsEndpoint {
        filter?: DeleteAttachmentsRequestFilter;
        hash?: string;
    }
    
    export interface DeleteAttachmentsEndpointFilter {
        and?: DeleteAttachmentsEndpointFilter[];
        hash?: string;
        not?: DeleteAttachmentsEndpointFilter[];
        or?: DeleteAttachmentsEndpointFilter[];
        set?: boolean;
    }
    
    export interface DeleteAttachmentsRequestFilter {
        and?: DeleteAttachmentsRequestFilter[];
        attachments?: AttachmentListFilter;
        hash?: string;
        ids?: IdListFilter;
        meta?: RequestMetaFilter;
        mode?: DeleteModeFilter;
        not?: DeleteAttachmentsRequestFilter[];
        or?: DeleteAttachmentsRequestFilter[];
        set?: boolean;
    }
    
    export interface DeleteAttachmentsResponseFilter {
        and?: DeleteAttachmentsResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: DeleteAttachmentsResponseFilter[];
        or?: DeleteAttachmentsResponseFilter[];
        set?: boolean;
    }
    
    export interface DeleteBlueWhateversEndpoint {
        filter?: DeleteBlueWhateversRequestFilter;
        hash?: string;
    }
    
    export interface DeleteBlueWhateversEndpointFilter {
        and?: DeleteBlueWhateversEndpointFilter[];
        hash?: string;
        not?: DeleteBlueWhateversEndpointFilter[];
        or?: DeleteBlueWhateversEndpointFilter[];
        set?: boolean;
    }
    
    export interface DeleteBlueWhateversRequestFilter {
        and?: DeleteBlueWhateversRequestFilter[];
        blueWhatevers?: BlueWhateverListFilter;
        hash?: string;
        ids?: IdListFilter;
        meta?: RequestMetaFilter;
        mode?: DeleteModeFilter;
        not?: DeleteBlueWhateversRequestFilter[];
        or?: DeleteBlueWhateversRequestFilter[];
        set?: boolean;
    }
    
    export interface DeleteBlueWhateversResponseFilter {
        and?: DeleteBlueWhateversResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: DeleteBlueWhateversResponseFilter[];
        or?: DeleteBlueWhateversResponseFilter[];
        set?: boolean;
    }
    
    export interface DeleteClientAccountsEndpoint {
        filter?: DeleteClientAccountsRequestFilter;
        hash?: string;
    }
    
    export interface DeleteClientAccountsEndpointFilter {
        and?: DeleteClientAccountsEndpointFilter[];
        hash?: string;
        not?: DeleteClientAccountsEndpointFilter[];
        or?: DeleteClientAccountsEndpointFilter[];
        set?: boolean;
    }
    
    export interface DeleteClientAccountsRequestFilter {
        and?: DeleteClientAccountsRequestFilter[];
        clientAccounts?: ClientAccountListFilter;
        hash?: string;
        ids?: IdListFilter;
        meta?: RequestMetaFilter;
        mode?: DeleteModeFilter;
        not?: DeleteClientAccountsRequestFilter[];
        or?: DeleteClientAccountsRequestFilter[];
        set?: boolean;
    }
    
    export interface DeleteClientAccountsResponseFilter {
        and?: DeleteClientAccountsResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: DeleteClientAccountsResponseFilter[];
        or?: DeleteClientAccountsResponseFilter[];
        set?: boolean;
    }
    
    export interface DeleteFeedsEndpoint {
        filter?: DeleteFeedsRequestFilter;
        hash?: string;
    }
    
    export interface DeleteFeedsEndpointFilter {
        and?: DeleteFeedsEndpointFilter[];
        hash?: string;
        not?: DeleteFeedsEndpointFilter[];
        or?: DeleteFeedsEndpointFilter[];
        set?: boolean;
    }
    
    export interface DeleteFeedsRequestFilter {
        and?: DeleteFeedsRequestFilter[];
        feeds?: FeedListFilter;
        hash?: string;
        ids?: IdListFilter;
        meta?: RequestMetaFilter;
        mode?: DeleteModeFilter;
        not?: DeleteFeedsRequestFilter[];
        or?: DeleteFeedsRequestFilter[];
        set?: boolean;
    }
    
    export interface DeleteFeedsResponseFilter {
        and?: DeleteFeedsResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: DeleteFeedsResponseFilter[];
        or?: DeleteFeedsResponseFilter[];
        set?: boolean;
    }
    
    export interface DeleteModeFilter {
        and?: DeleteModeFilter[];
        hash?: string;
        kind?: EnumFilter;
        not?: DeleteModeFilter[];
        or?: DeleteModeFilter[];
        set?: boolean;
    }
    
    export interface DeletePeopleEndpoint {
        filter?: DeletePeopleRequestFilter;
        hash?: string;
    }
    
    export interface DeletePeopleEndpointFilter {
        and?: DeletePeopleEndpointFilter[];
        hash?: string;
        not?: DeletePeopleEndpointFilter[];
        or?: DeletePeopleEndpointFilter[];
        set?: boolean;
    }
    
    export interface DeletePeopleRequestFilter {
        and?: DeletePeopleRequestFilter[];
        hash?: string;
        ids?: IdListFilter;
        meta?: RequestMetaFilter;
        mode?: DeleteModeFilter;
        not?: DeletePeopleRequestFilter[];
        or?: DeletePeopleRequestFilter[];
        people?: PersonListFilter;
        set?: boolean;
    }
    
    export interface DeletePeopleResponseFilter {
        and?: DeletePeopleResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: DeletePeopleResponseFilter[];
        or?: DeletePeopleResponseFilter[];
        set?: boolean;
    }
    
    export interface DeleteServiceAccountsEndpoint {
        filter?: DeleteServiceAccountsRequestFilter;
        hash?: string;
    }
    
    export interface DeleteServiceAccountsEndpointFilter {
        and?: DeleteServiceAccountsEndpointFilter[];
        hash?: string;
        not?: DeleteServiceAccountsEndpointFilter[];
        or?: DeleteServiceAccountsEndpointFilter[];
        set?: boolean;
    }
    
    export interface DeleteServiceAccountsRequestFilter {
        and?: DeleteServiceAccountsRequestFilter[];
        hash?: string;
        ids?: IdListFilter;
        meta?: RequestMetaFilter;
        mode?: DeleteModeFilter;
        not?: DeleteServiceAccountsRequestFilter[];
        or?: DeleteServiceAccountsRequestFilter[];
        serviceAccounts?: ServiceAccountListFilter;
        set?: boolean;
    }
    
    export interface DeleteServiceAccountsResponseFilter {
        and?: DeleteServiceAccountsResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: DeleteServiceAccountsResponseFilter[];
        or?: DeleteServiceAccountsResponseFilter[];
        set?: boolean;
    }
    
    export interface DeleteServicesEndpoint {
        filter?: DeleteServicesRequestFilter;
        hash?: string;
    }
    
    export interface DeleteServicesEndpointFilter {
        and?: DeleteServicesEndpointFilter[];
        hash?: string;
        not?: DeleteServicesEndpointFilter[];
        or?: DeleteServicesEndpointFilter[];
        set?: boolean;
    }
    
    export interface DeleteServicesRequestFilter {
        and?: DeleteServicesRequestFilter[];
        hash?: string;
        ids?: IdListFilter;
        meta?: RequestMetaFilter;
        mode?: DeleteModeFilter;
        not?: DeleteServicesRequestFilter[];
        or?: DeleteServicesRequestFilter[];
        services?: ServiceListFilter;
        set?: boolean;
    }
    
    export interface DeleteServicesResponseFilter {
        and?: DeleteServicesResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: DeleteServicesResponseFilter[];
        or?: DeleteServicesResponseFilter[];
        set?: boolean;
    }
    
    export interface DeleteStatusesEndpoint {
        filter?: DeleteStatusesRequestFilter;
        hash?: string;
    }
    
    export interface DeleteStatusesEndpointFilter {
        and?: DeleteStatusesEndpointFilter[];
        hash?: string;
        not?: DeleteStatusesEndpointFilter[];
        or?: DeleteStatusesEndpointFilter[];
        set?: boolean;
    }
    
    export interface DeleteStatusesRequestFilter {
        and?: DeleteStatusesRequestFilter[];
        hash?: string;
        ids?: IdListFilter;
        meta?: RequestMetaFilter;
        mode?: DeleteModeFilter;
        not?: DeleteStatusesRequestFilter[];
        or?: DeleteStatusesRequestFilter[];
        set?: boolean;
        statuses?: StatusListFilter;
    }
    
    export interface DeleteStatusesResponseFilter {
        and?: DeleteStatusesResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: DeleteStatusesResponseFilter[];
        or?: DeleteStatusesResponseFilter[];
        set?: boolean;
    }
    
    export interface DeleteWhateversEndpoint {
        filter?: DeleteWhateversRequestFilter;
        hash?: string;
    }
    
    export interface DeleteWhateversEndpointFilter {
        and?: DeleteWhateversEndpointFilter[];
        hash?: string;
        not?: DeleteWhateversEndpointFilter[];
        or?: DeleteWhateversEndpointFilter[];
        set?: boolean;
    }
    
    export interface DeleteWhateversRequestFilter {
        and?: DeleteWhateversRequestFilter[];
        hash?: string;
        ids?: IdListFilter;
        meta?: RequestMetaFilter;
        mode?: DeleteModeFilter;
        not?: DeleteWhateversRequestFilter[];
        or?: DeleteWhateversRequestFilter[];
        set?: boolean;
        whatevers?: WhateverListFilter;
    }
    
    export interface DeleteWhateversResponseFilter {
        and?: DeleteWhateversResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: DeleteWhateversResponseFilter[];
        or?: DeleteWhateversResponseFilter[];
        set?: boolean;
    }
    
    export interface Email {
        hash?: string;
        value?: string;
    }
    
    export interface EmailFilter {
        and?: EmailFilter[];
        hash?: string;
        not?: EmailFilter[];
        or?: EmailFilter[];
        set?: boolean;
        value?: StringFilter;
    }
    
    export interface EmailSelect {
        hash?: string;
        selectAll?: boolean;
        value?: boolean;
    }
    
    export interface Endpoints {
        authenticateClientAccount?: AuthenticateClientAccountEndpoint;
        deleteAttachments?: DeleteAttachmentsEndpoint;
        deleteBlueWhatevers?: DeleteBlueWhateversEndpoint;
        deleteClientAccounts?: DeleteClientAccountsEndpoint;
        deleteFeeds?: DeleteFeedsEndpoint;
        deletePeople?: DeletePeopleEndpoint;
        deleteServiceAccounts?: DeleteServiceAccountsEndpoint;
        deleteServices?: DeleteServicesEndpoint;
        deleteStatuses?: DeleteStatusesEndpoint;
        deleteWhatevers?: DeleteWhateversEndpoint;
        getAttachments?: GetAttachmentsEndpoint;
        getBlueWhatevers?: GetBlueWhateversEndpoint;
        getClientAccounts?: GetClientAccountsEndpoint;
        getFeeds?: GetFeedsEndpoint;
        getPeople?: GetPeopleEndpoint;
        getServiceAccounts?: GetServiceAccountsEndpoint;
        getServices?: GetServicesEndpoint;
        getStatuses?: GetStatusesEndpoint;
        getWhatevers?: GetWhateversEndpoint;
        hash?: string;
        lookupService?: LookupServiceEndpoint;
        pipeAttachments?: PipeAttachmentsEndpoint;
        pipeBlueWhatevers?: PipeBlueWhateversEndpoint;
        pipeClientAccounts?: PipeClientAccountsEndpoint;
        pipeFeeds?: PipeFeedsEndpoint;
        pipePeople?: PipePeopleEndpoint;
        pipeServiceAccounts?: PipeServiceAccountsEndpoint;
        pipeServices?: PipeServicesEndpoint;
        pipeStatuses?: PipeStatusesEndpoint;
        pipeWhatevers?: PipeWhateversEndpoint;
        postAttachments?: PostAttachmentsEndpoint;
        postBlueWhatevers?: PostBlueWhateversEndpoint;
        postClientAccounts?: PostClientAccountsEndpoint;
        postFeeds?: PostFeedsEndpoint;
        postPeople?: PostPeopleEndpoint;
        postServiceAccounts?: PostServiceAccountsEndpoint;
        postServices?: PostServicesEndpoint;
        postStatuses?: PostStatusesEndpoint;
        postWhatevers?: PostWhateversEndpoint;
        putAttachments?: PutAttachmentsEndpoint;
        putBlueWhatevers?: PutBlueWhateversEndpoint;
        putClientAccounts?: PutClientAccountsEndpoint;
        putFeeds?: PutFeedsEndpoint;
        putPeople?: PutPeopleEndpoint;
        putServiceAccounts?: PutServiceAccountsEndpoint;
        putServices?: PutServicesEndpoint;
        putStatuses?: PutStatusesEndpoint;
        putWhatevers?: PutWhateversEndpoint;
        verifyToken?: VerifyTokenEndpoint;
    }
    
    export interface EndpointsFilter {
        and?: EndpointsFilter[];
        authenticateClientAccount?: AuthenticateClientAccountEndpointFilter;
        deleteAttachments?: DeleteAttachmentsEndpointFilter;
        deleteBlueWhatevers?: DeleteBlueWhateversEndpointFilter;
        deleteClientAccounts?: DeleteClientAccountsEndpointFilter;
        deleteFeeds?: DeleteFeedsEndpointFilter;
        deletePeople?: DeletePeopleEndpointFilter;
        deleteServiceAccounts?: DeleteServiceAccountsEndpointFilter;
        deleteServices?: DeleteServicesEndpointFilter;
        deleteStatuses?: DeleteStatusesEndpointFilter;
        deleteWhatevers?: DeleteWhateversEndpointFilter;
        getAttachments?: GetAttachmentsEndpointFilter;
        getBlueWhatevers?: GetBlueWhateversEndpointFilter;
        getClientAccounts?: GetClientAccountsEndpointFilter;
        getFeeds?: GetFeedsEndpointFilter;
        getPeople?: GetPeopleEndpointFilter;
        getServiceAccounts?: GetServiceAccountsEndpointFilter;
        getServices?: GetServicesEndpointFilter;
        getStatuses?: GetStatusesEndpointFilter;
        getWhatevers?: GetWhateversEndpointFilter;
        hash?: string;
        lookupService?: LookupServiceEndpointFilter;
        not?: EndpointsFilter[];
        or?: EndpointsFilter[];
        pipeAttachments?: PipeAttachmentsEndpointFilter;
        pipeBlueWhatevers?: PipeBlueWhateversEndpointFilter;
        pipeClientAccounts?: PipeClientAccountsEndpointFilter;
        pipeFeeds?: PipeFeedsEndpointFilter;
        pipePeople?: PipePeopleEndpointFilter;
        pipeServiceAccounts?: PipeServiceAccountsEndpointFilter;
        pipeServices?: PipeServicesEndpointFilter;
        pipeStatuses?: PipeStatusesEndpointFilter;
        pipeWhatevers?: PipeWhateversEndpointFilter;
        postAttachments?: PostAttachmentsEndpointFilter;
        postBlueWhatevers?: PostBlueWhateversEndpointFilter;
        postClientAccounts?: PostClientAccountsEndpointFilter;
        postFeeds?: PostFeedsEndpointFilter;
        postPeople?: PostPeopleEndpointFilter;
        postServiceAccounts?: PostServiceAccountsEndpointFilter;
        postServices?: PostServicesEndpointFilter;
        postStatuses?: PostStatusesEndpointFilter;
        postWhatevers?: PostWhateversEndpointFilter;
        putAttachments?: PutAttachmentsEndpointFilter;
        putBlueWhatevers?: PutBlueWhateversEndpointFilter;
        putClientAccounts?: PutClientAccountsEndpointFilter;
        putFeeds?: PutFeedsEndpointFilter;
        putPeople?: PutPeopleEndpointFilter;
        putServiceAccounts?: PutServiceAccountsEndpointFilter;
        putServices?: PutServicesEndpointFilter;
        putStatuses?: PutStatusesEndpointFilter;
        putWhatevers?: PutWhateversEndpointFilter;
        set?: boolean;
        verifyToken?: VerifyTokenEndpointFilter;
    }
    
    export interface EnumFilter {
        and?: EnumFilter[];
        hash?: string;
        in?: string[];
        is?: string;
        not?: string;
        notIn?: string[];
        or?: EnumFilter[];
        set?: boolean;
    }
    
    export interface EnumListFilter {
        and?: EnumFilter;
        hash?: string;
        not?: EnumFilter;
        or?: EnumFilter;
    }
    
    export interface Error {
        hash?: string;
        id?: Id;
        kind?: string;
        message?: Text;
        service?: Service;
        wraps?: Error;
    }
    
    export interface ErrorFilter {
        and?: ErrorFilter[];
        hash?: string;
        id?: IdFilter;
        kind?: EnumFilter;
        message?: TextFilter;
        not?: ErrorFilter[];
        or?: ErrorFilter[];
        service?: ServiceFilter;
        set?: boolean;
        wraps?: ErrorFilter;
    }
    
    export interface ErrorListFilter {
        every?: ErrorFilter;
        hash?: string;
        none?: ErrorFilter;
        some?: ErrorFilter;
    }
    
    export interface ErrorSelect {
        hash?: string;
        id?: IdSelect;
        kind?: boolean;
        message?: TextSelect;
        selectAll?: boolean;
        service?: ServiceSelect;
        wraps?: ErrorSelect;
    }
    
    export interface FeedFilter {
        alternativeIds?: IdListFilter;
        and?: FeedFilter[];
        hash?: string;
        id?: ServiceIdFilter;
        info?: InfoFilter;
        kind?: EnumFilter;
        meta?: TypeMetaFilter;
        not?: FeedFilter[];
        or?: FeedFilter[];
        set?: boolean;
    }
    
    export interface FeedListFilter {
        every?: FeedFilter;
        hash?: string;
        none?: FeedFilter;
        some?: FeedFilter;
    }
    
    export interface Float64Filter {
        and?: Float64Filter[];
        gt?: number;
        gte?: number;
        hash?: string;
        in?: number[];
        is?: number;
        lt?: number;
        lte?: number;
        not?: number;
        notIn?: number[];
        or?: Float64Filter[];
        set?: boolean;
    }
    
    export interface Float64ListFilter {
        and?: Float64Filter;
        hash?: string;
        not?: Float64Filter;
        or?: Float64Filter;
    }
    
    export interface FloatRange {
        from?: number;
        hash?: string;
        to?: number;
    }
    
    export interface FloatRangeFilter {
        and?: FloatRangeFilter[];
        from?: Float64Filter;
        hash?: string;
        not?: FloatRangeFilter[];
        or?: FloatRangeFilter[];
        set?: boolean;
        to?: Float64Filter;
    }
    
    export interface GetAttachmentsEndpoint {
        filter?: GetAttachmentsRequestFilter;
        hash?: string;
    }
    
    export interface GetAttachmentsEndpointFilter {
        and?: GetAttachmentsEndpointFilter[];
        hash?: string;
        not?: GetAttachmentsEndpointFilter[];
        or?: GetAttachmentsEndpointFilter[];
        set?: boolean;
    }
    
    export interface GetAttachmentsRequestFilter {
        and?: GetAttachmentsRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: GetModeFilter;
        not?: GetAttachmentsRequestFilter[];
        or?: GetAttachmentsRequestFilter[];
        pages?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface GetAttachmentsResponseFilter {
        and?: GetAttachmentsResponseFilter[];
        attachments?: AttachmentListFilter;
        hash?: string;
        meta?: CollectionMetaFilter;
        not?: GetAttachmentsResponseFilter[];
        or?: GetAttachmentsResponseFilter[];
        set?: boolean;
    }
    
    export interface GetBlueWhateversCollection {
        filter?: BlueWhateverFilter;
        hash?: string;
        pages?: ServicePage[];
        relations?: GetBlueWhateversRelations;
        select?: BlueWhateversCollectionSelect;
        serviceFilter?: ServiceFilter;
        sort?: BlueWhateverSort;
    }
    
    export interface GetBlueWhateversEndpoint {
        filter?: GetBlueWhateversRequestFilter;
        hash?: string;
    }
    
    export interface GetBlueWhateversEndpointFilter {
        and?: GetBlueWhateversEndpointFilter[];
        hash?: string;
        not?: GetBlueWhateversEndpointFilter[];
        or?: GetBlueWhateversEndpointFilter[];
        set?: boolean;
    }
    
    export interface GetBlueWhateversRelations {
        hash?: string;
        knewByWhatevers?: GetWhateversCollection;
    }
    
    export interface GetBlueWhateversRequestFilter {
        and?: GetBlueWhateversRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: GetModeFilter;
        not?: GetBlueWhateversRequestFilter[];
        or?: GetBlueWhateversRequestFilter[];
        pages?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface GetBlueWhateversResponseFilter {
        and?: GetBlueWhateversResponseFilter[];
        blueWhatevers?: BlueWhateverListFilter;
        hash?: string;
        meta?: CollectionMetaFilter;
        not?: GetBlueWhateversResponseFilter[];
        or?: GetBlueWhateversResponseFilter[];
        set?: boolean;
    }
    
    export interface GetClientAccountsCollection {
        filter?: ClientAccountFilter;
        hash?: string;
        pages?: ServicePage[];
        relations?: GetClientAccountsRelations;
        select?: ClientAccountsCollectionSelect;
        serviceFilter?: ServiceFilter;
        sort?: ClientAccountSort;
    }
    
    export interface GetClientAccountsEndpoint {
        filter?: GetClientAccountsRequestFilter;
        hash?: string;
    }
    
    export interface GetClientAccountsEndpointFilter {
        and?: GetClientAccountsEndpointFilter[];
        hash?: string;
        not?: GetClientAccountsEndpointFilter[];
        or?: GetClientAccountsEndpointFilter[];
        set?: boolean;
    }
    
    export interface GetClientAccountsRelations {
        hash?: string;
        ownsServiceAccounts?: GetServiceAccountsCollection;
    }
    
    export interface GetClientAccountsRequest {
        auth?: Auth;
        filter?: ClientAccountFilter;
        hash?: string;
        meta?: RequestMeta;
        mode?: GetMode;
        pages?: ServicePage[];
        relations?: GetClientAccountsRelations;
        select?: GetClientAccountsResponseSelect;
        serviceFilter?: ServiceFilter;
        sort?: ClientAccountSort;
    }
    
    export interface GetClientAccountsRequestFilter {
        and?: GetClientAccountsRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: GetModeFilter;
        not?: GetClientAccountsRequestFilter[];
        or?: GetClientAccountsRequestFilter[];
        pages?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface GetClientAccountsResponse {
        clientAccounts?: ClientAccount[];
        hash?: string;
        meta?: CollectionMeta;
    }
    
    export interface GetClientAccountsResponseFilter {
        and?: GetClientAccountsResponseFilter[];
        clientAccounts?: ClientAccountListFilter;
        hash?: string;
        meta?: CollectionMetaFilter;
        not?: GetClientAccountsResponseFilter[];
        or?: GetClientAccountsResponseFilter[];
        set?: boolean;
    }
    
    export interface GetClientAccountsResponseSelect {
        clientAccounts?: ClientAccountSelect;
        hash?: string;
        meta?: CollectionMetaSelect;
        selectAll?: boolean;
    }
    
    export interface GetFeedsEndpoint {
        filter?: GetFeedsRequestFilter;
        hash?: string;
    }
    
    export interface GetFeedsEndpointFilter {
        and?: GetFeedsEndpointFilter[];
        hash?: string;
        not?: GetFeedsEndpointFilter[];
        or?: GetFeedsEndpointFilter[];
        set?: boolean;
    }
    
    export interface GetFeedsRequestFilter {
        and?: GetFeedsRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: GetModeFilter;
        not?: GetFeedsRequestFilter[];
        or?: GetFeedsRequestFilter[];
        pages?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface GetFeedsResponseFilter {
        and?: GetFeedsResponseFilter[];
        feeds?: FeedListFilter;
        hash?: string;
        meta?: CollectionMetaFilter;
        not?: GetFeedsResponseFilter[];
        or?: GetFeedsResponseFilter[];
        set?: boolean;
    }
    
    export interface GetMode {
        collection?: CollectionGetMode;
        hash?: string;
        id?: Id;
        kind?: string;
        relation?: RelationGetMode;
        search?: SearchGetMode;
    }
    
    export interface GetModeFilter {
        and?: GetModeFilter[];
        collection?: CollectionGetModeFilter;
        hash?: string;
        id?: IdFilter;
        kind?: EnumFilter;
        not?: GetModeFilter[];
        or?: GetModeFilter[];
        relation?: RelationGetModeFilter;
        search?: SearchGetModeFilter;
        set?: boolean;
    }
    
    export interface GetPeopleEndpoint {
        filter?: GetPeopleRequestFilter;
        hash?: string;
    }
    
    export interface GetPeopleEndpointFilter {
        and?: GetPeopleEndpointFilter[];
        hash?: string;
        not?: GetPeopleEndpointFilter[];
        or?: GetPeopleEndpointFilter[];
        set?: boolean;
    }
    
    export interface GetPeopleRequestFilter {
        and?: GetPeopleRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: GetModeFilter;
        not?: GetPeopleRequestFilter[];
        or?: GetPeopleRequestFilter[];
        pages?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface GetPeopleResponseFilter {
        and?: GetPeopleResponseFilter[];
        hash?: string;
        meta?: CollectionMetaFilter;
        not?: GetPeopleResponseFilter[];
        or?: GetPeopleResponseFilter[];
        people?: PersonListFilter;
        set?: boolean;
    }
    
    export interface GetServiceAccountsCollection {
        filter?: ServiceAccountFilter;
        hash?: string;
        pages?: ServicePage[];
        relations?: GetServiceAccountsRelations;
        select?: ServiceAccountsCollectionSelect;
        serviceFilter?: ServiceFilter;
        sort?: ServiceAccountSort;
    }
    
    export interface GetServiceAccountsEndpoint {
        filter?: GetServiceAccountsRequestFilter;
        hash?: string;
    }
    
    export interface GetServiceAccountsEndpointFilter {
        and?: GetServiceAccountsEndpointFilter[];
        hash?: string;
        not?: GetServiceAccountsEndpointFilter[];
        or?: GetServiceAccountsEndpointFilter[];
        set?: boolean;
    }
    
    export interface GetServiceAccountsRelations {
        hash?: string;
        ownedByClientAccounts?: GetClientAccountsCollection;
        usedByServices?: GetServicesCollection;
    }
    
    export interface GetServiceAccountsRequest {
        auth?: Auth;
        filter?: ServiceAccountFilter;
        hash?: string;
        meta?: RequestMeta;
        mode?: GetMode;
        pages?: ServicePage[];
        relations?: GetServiceAccountsRelations;
        select?: GetServiceAccountsResponseSelect;
        serviceFilter?: ServiceFilter;
        sort?: ServiceAccountSort;
    }
    
    export interface GetServiceAccountsRequestFilter {
        and?: GetServiceAccountsRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: GetModeFilter;
        not?: GetServiceAccountsRequestFilter[];
        or?: GetServiceAccountsRequestFilter[];
        pages?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface GetServiceAccountsResponse {
        hash?: string;
        meta?: CollectionMeta;
        serviceAccounts?: ServiceAccount[];
    }
    
    export interface GetServiceAccountsResponseFilter {
        and?: GetServiceAccountsResponseFilter[];
        hash?: string;
        meta?: CollectionMetaFilter;
        not?: GetServiceAccountsResponseFilter[];
        or?: GetServiceAccountsResponseFilter[];
        serviceAccounts?: ServiceAccountListFilter;
        set?: boolean;
    }
    
    export interface GetServiceAccountsResponseSelect {
        hash?: string;
        meta?: CollectionMetaSelect;
        selectAll?: boolean;
        serviceAccounts?: ServiceAccountSelect;
    }
    
    export interface GetServicesCollection {
        filter?: ServiceFilter;
        hash?: string;
        pages?: ServicePage[];
        relations?: GetServicesRelations;
        select?: ServicesCollectionSelect;
        serviceFilter?: ServiceFilter;
        sort?: ServiceSort;
    }
    
    export interface GetServicesEndpoint {
        filter?: GetServicesRequestFilter;
        hash?: string;
    }
    
    export interface GetServicesEndpointFilter {
        and?: GetServicesEndpointFilter[];
        hash?: string;
        not?: GetServicesEndpointFilter[];
        or?: GetServicesEndpointFilter[];
        set?: boolean;
    }
    
    export interface GetServicesRelations {
        hash?: string;
        usesServiceAccounts?: GetServiceAccountsCollection;
    }
    
    export interface GetServicesRequest {
        auth?: Auth;
        filter?: ServiceFilter;
        hash?: string;
        meta?: RequestMeta;
        mode?: GetMode;
        pages?: ServicePage[];
        relations?: GetServicesRelations;
        select?: GetServicesResponseSelect;
        serviceFilter?: ServiceFilter;
        sort?: ServiceSort;
    }
    
    export interface GetServicesRequestFilter {
        and?: GetServicesRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: GetModeFilter;
        not?: GetServicesRequestFilter[];
        or?: GetServicesRequestFilter[];
        pages?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface GetServicesResponse {
        hash?: string;
        meta?: CollectionMeta;
        services?: Service[];
    }
    
    export interface GetServicesResponseFilter {
        and?: GetServicesResponseFilter[];
        hash?: string;
        meta?: CollectionMetaFilter;
        not?: GetServicesResponseFilter[];
        or?: GetServicesResponseFilter[];
        services?: ServiceListFilter;
        set?: boolean;
    }
    
    export interface GetServicesResponseSelect {
        hash?: string;
        meta?: CollectionMetaSelect;
        selectAll?: boolean;
        services?: ServiceSelect;
    }
    
    export interface GetStatusesEndpoint {
        filter?: GetStatusesRequestFilter;
        hash?: string;
    }
    
    export interface GetStatusesEndpointFilter {
        and?: GetStatusesEndpointFilter[];
        hash?: string;
        not?: GetStatusesEndpointFilter[];
        or?: GetStatusesEndpointFilter[];
        set?: boolean;
    }
    
    export interface GetStatusesRequestFilter {
        and?: GetStatusesRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: GetModeFilter;
        not?: GetStatusesRequestFilter[];
        or?: GetStatusesRequestFilter[];
        pages?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface GetStatusesResponseFilter {
        and?: GetStatusesResponseFilter[];
        hash?: string;
        meta?: CollectionMetaFilter;
        not?: GetStatusesResponseFilter[];
        or?: GetStatusesResponseFilter[];
        set?: boolean;
        statuses?: StatusListFilter;
    }
    
    export interface GetWhateversCollection {
        filter?: WhateverFilter;
        hash?: string;
        pages?: ServicePage[];
        relations?: GetWhateversRelations;
        select?: WhateversCollectionSelect;
        serviceFilter?: ServiceFilter;
        sort?: WhateverSort;
    }
    
    export interface GetWhateversEndpoint {
        filter?: GetWhateversRequestFilter;
        hash?: string;
    }
    
    export interface GetWhateversEndpointFilter {
        and?: GetWhateversEndpointFilter[];
        hash?: string;
        not?: GetWhateversEndpointFilter[];
        or?: GetWhateversEndpointFilter[];
        set?: boolean;
    }
    
    export interface GetWhateversRelations {
        hash?: string;
        knewByWhatevers?: GetWhateversCollection;
        knowsBlueWhatevers?: GetBlueWhateversCollection;
        knowsWhatevers?: GetWhateversCollection;
    }
    
    export interface GetWhateversRequest {
        auth?: Auth;
        filter?: WhateverFilter;
        hash?: string;
        meta?: RequestMeta;
        mode?: GetMode;
        pages?: ServicePage[];
        relations?: GetWhateversRelations;
        select?: GetWhateversResponseSelect;
        serviceFilter?: ServiceFilter;
        sort?: WhateverSort;
    }
    
    export interface GetWhateversRequestFilter {
        and?: GetWhateversRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: GetModeFilter;
        not?: GetWhateversRequestFilter[];
        or?: GetWhateversRequestFilter[];
        pages?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface GetWhateversResponse {
        hash?: string;
        meta?: CollectionMeta;
        whatevers?: Whatever[];
    }
    
    export interface GetWhateversResponseFilter {
        and?: GetWhateversResponseFilter[];
        hash?: string;
        meta?: CollectionMetaFilter;
        not?: GetWhateversResponseFilter[];
        or?: GetWhateversResponseFilter[];
        set?: boolean;
        whatevers?: WhateverListFilter;
    }
    
    export interface GetWhateversResponseSelect {
        hash?: string;
        meta?: CollectionMetaSelect;
        selectAll?: boolean;
        whatevers?: WhateverSelect;
    }
    
    export interface Id {
        ean?: string;
        email?: Email;
        hash?: string;
        kind?: string;
        local?: string;
        me?: boolean;
        name?: string;
        serviceId?: ServiceId;
        token?: Token;
        url?: Url;
        username?: string;
    }
    
    export interface IdFilter {
        and?: IdFilter[];
        ean?: StringFilter;
        email?: EmailFilter;
        hash?: string;
        kind?: EnumFilter;
        local?: StringFilter;
        me?: BoolFilter;
        name?: StringFilter;
        not?: IdFilter[];
        or?: IdFilter[];
        serviceId?: ServiceIdFilter;
        set?: boolean;
        token?: TokenFilter;
        url?: UrlFilter;
        username?: StringFilter;
    }
    
    export interface IdListFilter {
        every?: IdFilter;
        hash?: string;
        none?: IdFilter;
        some?: IdFilter;
    }
    
    export interface IdSelect {
        ean?: boolean;
        email?: EmailSelect;
        hash?: string;
        kind?: boolean;
        local?: boolean;
        me?: boolean;
        name?: boolean;
        selectAll?: boolean;
        serviceId?: ServiceIdSelect;
        token?: TokenSelect;
        url?: UrlSelect;
        username?: boolean;
    }
    
    export interface ImageFilter {
        and?: ImageFilter[];
        description?: TextFilter;
        hash?: string;
        height?: Int32Filter;
        isPreview?: BoolFilter;
        not?: ImageFilter[];
        or?: ImageFilter[];
        set?: boolean;
        url?: UrlFilter;
        width?: Int32Filter;
    }
    
    export interface IndexPage {
        hash?: string;
        page?: number;
    }
    
    export interface IndexPageFilter {
        and?: IndexPageFilter[];
        hash?: string;
        not?: IndexPageFilter[];
        or?: IndexPageFilter[];
        page?: Int32Filter;
        set?: boolean;
    }
    
    export interface IndexPageSelect {
        hash?: string;
        page?: boolean;
        selectAll?: boolean;
    }
    
    export interface InfoFilter {
        and?: InfoFilter[];
        description?: TextFilter;
        hash?: string;
        name?: TextFilter;
        not?: InfoFilter[];
        or?: InfoFilter[];
        purpose?: TextFilter;
        set?: boolean;
    }
    
    export interface Int32Filter {
        and?: Int32Filter[];
        gt?: number;
        gte?: number;
        hash?: string;
        in?: number[];
        is?: number;
        lt?: number;
        lte?: number;
        not?: number;
        notIn?: number[];
        or?: Int32Filter[];
        set?: boolean;
    }
    
    export interface Int32ListFilter {
        and?: Int32Filter;
        hash?: string;
        not?: Int32Filter;
        or?: Int32Filter;
    }
    
    export interface LengthValue {
        hash?: string;
        isEstimate?: boolean;
        kind?: string;
        range?: FloatRange;
        unit?: string;
        value?: number;
    }
    
    export interface LengthValueFilter {
        and?: LengthValueFilter[];
        hash?: string;
        isEstimate?: BoolFilter;
        kind?: EnumFilter;
        not?: LengthValueFilter[];
        or?: LengthValueFilter[];
        range?: FloatRangeFilter;
        set?: boolean;
        unit?: EnumFilter;
        value?: Float64Filter;
    }
    
    export interface LocationQuery {
        city?: string;
        cityDistrict?: string;
        country?: string;
        countryState?: string;
        countryStateDistrict?: string;
        hash?: string;
        radiusLt?: LengthValue;
        street?: string;
        zipCode?: string;
    }
    
    export interface LocationQueryFilter {
        and?: LocationQueryFilter[];
        city?: StringFilter;
        cityDistrict?: StringFilter;
        country?: StringFilter;
        countryState?: StringFilter;
        countryStateDistrict?: StringFilter;
        hash?: string;
        not?: LocationQueryFilter[];
        or?: LocationQueryFilter[];
        radiusLt?: LengthValueFilter;
        set?: boolean;
        street?: StringFilter;
        zipCode?: StringFilter;
    }
    
    export interface LookupServiceEndpoint {
        filter?: LookupServiceRequestFilter;
        hash?: string;
    }
    
    export interface LookupServiceEndpointFilter {
        and?: LookupServiceEndpointFilter[];
        hash?: string;
        not?: LookupServiceEndpointFilter[];
        or?: LookupServiceEndpointFilter[];
        set?: boolean;
    }
    
    export interface LookupServiceRequestFilter {
        and?: LookupServiceRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        not?: LookupServiceRequestFilter[];
        or?: LookupServiceRequestFilter[];
        set?: boolean;
    }
    
    export interface OffsetPage {
        hash?: string;
        limit?: number;
        offset?: number;
    }
    
    export interface OffsetPageFilter {
        and?: OffsetPageFilter[];
        hash?: string;
        limit?: Int32Filter;
        not?: OffsetPageFilter[];
        offset?: Int32Filter;
        or?: OffsetPageFilter[];
        set?: boolean;
    }
    
    export interface OffsetPageSelect {
        hash?: string;
        limit?: boolean;
        offset?: boolean;
        selectAll?: boolean;
    }
    
    export interface Page {
        cursorPage?: CursorPage;
        hash?: string;
        indexPage?: IndexPage;
        kind?: string;
        offsetPage?: OffsetPage;
    }
    
    export interface PageFilter {
        and?: PageFilter[];
        cursorPage?: CursorPageFilter;
        hash?: string;
        indexPage?: IndexPageFilter;
        kind?: EnumFilter;
        not?: PageFilter[];
        offsetPage?: OffsetPageFilter;
        or?: PageFilter[];
        set?: boolean;
    }
    
    export interface PageSelect {
        cursorPage?: CursorPageSelect;
        hash?: string;
        indexPage?: IndexPageSelect;
        kind?: boolean;
        offsetPage?: OffsetPageSelect;
        selectAll?: boolean;
    }
    
    export interface Pagination {
        current?: ServicePage[];
        hash?: string;
        next?: ServicePage[];
        previous?: ServicePage[];
    }
    
    export interface PaginationFilter {
        and?: PaginationFilter[];
        current?: ServicePageListFilter;
        hash?: string;
        next?: ServicePageListFilter;
        not?: PaginationFilter[];
        or?: PaginationFilter[];
        previous?: ServicePageListFilter;
        set?: boolean;
    }
    
    export interface PaginationSelect {
        current?: ServicePageSelect;
        hash?: string;
        next?: ServicePageSelect;
        previous?: ServicePageSelect;
        selectAll?: boolean;
    }
    
    export interface Password {
        hash?: string;
        hashFunction?: string;
        isHashed?: boolean;
        value?: string;
    }
    
    export interface PasswordFilter {
        and?: PasswordFilter[];
        hash?: string;
        hashFunction?: EnumFilter;
        isHashed?: BoolFilter;
        not?: PasswordFilter[];
        or?: PasswordFilter[];
        set?: boolean;
        value?: StringFilter;
    }
    
    export interface PasswordSelect {
        hash?: string;
        hashFunction?: boolean;
        isHashed?: boolean;
        selectAll?: boolean;
        value?: boolean;
    }
    
    export interface PasswordSort {
        hash?: string;
        isHashed?: string;
        value?: string;
    }
    
    export interface PersonFilter {
        alternativeIds?: IdListFilter;
        and?: PersonFilter[];
        avatar?: ImageFilter;
        displayName?: TextFilter;
        hash?: string;
        header?: ImageFilter;
        id?: ServiceIdFilter;
        meta?: TypeMetaFilter;
        not?: PersonFilter[];
        note?: TextFilter;
        or?: PersonFilter[];
        set?: boolean;
        username?: TextFilter;
    }
    
    export interface PersonListFilter {
        every?: PersonFilter;
        hash?: string;
        none?: PersonFilter;
        some?: PersonFilter;
    }
    
    export interface PipeAttachmentsContextFilter {
        and?: PipeAttachmentsContextFilter[];
        delete?: PipeDeleteAttachmentsContextFilter;
        get?: PipeGetAttachmentsContextFilter;
        hash?: string;
        not?: PipeAttachmentsContextFilter[];
        or?: PipeAttachmentsContextFilter[];
        post?: PipePostAttachmentsContextFilter;
        put?: PipePutAttachmentsContextFilter;
        set?: boolean;
    }
    
    export interface PipeAttachmentsEndpoint {
        filter?: PipeAttachmentsRequestFilter;
        hash?: string;
    }
    
    export interface PipeAttachmentsEndpointFilter {
        and?: PipeAttachmentsEndpointFilter[];
        hash?: string;
        not?: PipeAttachmentsEndpointFilter[];
        or?: PipeAttachmentsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PipeAttachmentsRequestFilter {
        and?: PipeAttachmentsRequestFilter[];
        context?: PipeAttachmentsContextFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PipeModeFilter;
        not?: PipeAttachmentsRequestFilter[];
        or?: PipeAttachmentsRequestFilter[];
        set?: boolean;
    }
    
    export interface PipeBlueWhateversContextFilter {
        and?: PipeBlueWhateversContextFilter[];
        delete?: PipeDeleteBlueWhateversContextFilter;
        get?: PipeGetBlueWhateversContextFilter;
        hash?: string;
        not?: PipeBlueWhateversContextFilter[];
        or?: PipeBlueWhateversContextFilter[];
        post?: PipePostBlueWhateversContextFilter;
        put?: PipePutBlueWhateversContextFilter;
        set?: boolean;
    }
    
    export interface PipeBlueWhateversEndpoint {
        filter?: PipeBlueWhateversRequestFilter;
        hash?: string;
    }
    
    export interface PipeBlueWhateversEndpointFilter {
        and?: PipeBlueWhateversEndpointFilter[];
        hash?: string;
        not?: PipeBlueWhateversEndpointFilter[];
        or?: PipeBlueWhateversEndpointFilter[];
        set?: boolean;
    }
    
    export interface PipeBlueWhateversRequestFilter {
        and?: PipeBlueWhateversRequestFilter[];
        context?: PipeBlueWhateversContextFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PipeModeFilter;
        not?: PipeBlueWhateversRequestFilter[];
        or?: PipeBlueWhateversRequestFilter[];
        set?: boolean;
    }
    
    export interface PipeClientAccountsContextFilter {
        and?: PipeClientAccountsContextFilter[];
        delete?: PipeDeleteClientAccountsContextFilter;
        get?: PipeGetClientAccountsContextFilter;
        hash?: string;
        not?: PipeClientAccountsContextFilter[];
        or?: PipeClientAccountsContextFilter[];
        post?: PipePostClientAccountsContextFilter;
        put?: PipePutClientAccountsContextFilter;
        set?: boolean;
    }
    
    export interface PipeClientAccountsEndpoint {
        filter?: PipeClientAccountsRequestFilter;
        hash?: string;
    }
    
    export interface PipeClientAccountsEndpointFilter {
        and?: PipeClientAccountsEndpointFilter[];
        hash?: string;
        not?: PipeClientAccountsEndpointFilter[];
        or?: PipeClientAccountsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PipeClientAccountsRequestFilter {
        and?: PipeClientAccountsRequestFilter[];
        context?: PipeClientAccountsContextFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PipeModeFilter;
        not?: PipeClientAccountsRequestFilter[];
        or?: PipeClientAccountsRequestFilter[];
        set?: boolean;
    }
    
    export interface PipeDeleteAttachmentsContextFilter {
        and?: PipeDeleteAttachmentsContextFilter[];
        clientRequest?: DeleteAttachmentsRequestFilter;
        clientResponse?: DeleteAttachmentsResponseFilter;
        hash?: string;
        not?: PipeDeleteAttachmentsContextFilter[];
        or?: PipeDeleteAttachmentsContextFilter[];
        serviceRequest?: DeleteAttachmentsRequestFilter;
        serviceResponse?: DeleteAttachmentsResponseFilter;
        set?: boolean;
    }
    
    export interface PipeDeleteBlueWhateversContextFilter {
        and?: PipeDeleteBlueWhateversContextFilter[];
        clientRequest?: DeleteBlueWhateversRequestFilter;
        clientResponse?: DeleteBlueWhateversResponseFilter;
        hash?: string;
        not?: PipeDeleteBlueWhateversContextFilter[];
        or?: PipeDeleteBlueWhateversContextFilter[];
        serviceRequest?: DeleteBlueWhateversRequestFilter;
        serviceResponse?: DeleteBlueWhateversResponseFilter;
        set?: boolean;
    }
    
    export interface PipeDeleteClientAccountsContextFilter {
        and?: PipeDeleteClientAccountsContextFilter[];
        clientRequest?: DeleteClientAccountsRequestFilter;
        clientResponse?: DeleteClientAccountsResponseFilter;
        hash?: string;
        not?: PipeDeleteClientAccountsContextFilter[];
        or?: PipeDeleteClientAccountsContextFilter[];
        serviceRequest?: DeleteClientAccountsRequestFilter;
        serviceResponse?: DeleteClientAccountsResponseFilter;
        set?: boolean;
    }
    
    export interface PipeDeleteFeedsContextFilter {
        and?: PipeDeleteFeedsContextFilter[];
        clientRequest?: DeleteFeedsRequestFilter;
        clientResponse?: DeleteFeedsResponseFilter;
        hash?: string;
        not?: PipeDeleteFeedsContextFilter[];
        or?: PipeDeleteFeedsContextFilter[];
        serviceRequest?: DeleteFeedsRequestFilter;
        serviceResponse?: DeleteFeedsResponseFilter;
        set?: boolean;
    }
    
    export interface PipeDeletePeopleContextFilter {
        and?: PipeDeletePeopleContextFilter[];
        clientRequest?: DeletePeopleRequestFilter;
        clientResponse?: DeletePeopleResponseFilter;
        hash?: string;
        not?: PipeDeletePeopleContextFilter[];
        or?: PipeDeletePeopleContextFilter[];
        serviceRequest?: DeletePeopleRequestFilter;
        serviceResponse?: DeletePeopleResponseFilter;
        set?: boolean;
    }
    
    export interface PipeDeleteServiceAccountsContextFilter {
        and?: PipeDeleteServiceAccountsContextFilter[];
        clientRequest?: DeleteServiceAccountsRequestFilter;
        clientResponse?: DeleteServiceAccountsResponseFilter;
        hash?: string;
        not?: PipeDeleteServiceAccountsContextFilter[];
        or?: PipeDeleteServiceAccountsContextFilter[];
        serviceRequest?: DeleteServiceAccountsRequestFilter;
        serviceResponse?: DeleteServiceAccountsResponseFilter;
        set?: boolean;
    }
    
    export interface PipeDeleteServicesContextFilter {
        and?: PipeDeleteServicesContextFilter[];
        clientRequest?: DeleteServicesRequestFilter;
        clientResponse?: DeleteServicesResponseFilter;
        hash?: string;
        not?: PipeDeleteServicesContextFilter[];
        or?: PipeDeleteServicesContextFilter[];
        serviceRequest?: DeleteServicesRequestFilter;
        serviceResponse?: DeleteServicesResponseFilter;
        set?: boolean;
    }
    
    export interface PipeDeleteStatusesContextFilter {
        and?: PipeDeleteStatusesContextFilter[];
        clientRequest?: DeleteStatusesRequestFilter;
        clientResponse?: DeleteStatusesResponseFilter;
        hash?: string;
        not?: PipeDeleteStatusesContextFilter[];
        or?: PipeDeleteStatusesContextFilter[];
        serviceRequest?: DeleteStatusesRequestFilter;
        serviceResponse?: DeleteStatusesResponseFilter;
        set?: boolean;
    }
    
    export interface PipeDeleteWhateversContextFilter {
        and?: PipeDeleteWhateversContextFilter[];
        clientRequest?: DeleteWhateversRequestFilter;
        clientResponse?: DeleteWhateversResponseFilter;
        hash?: string;
        not?: PipeDeleteWhateversContextFilter[];
        or?: PipeDeleteWhateversContextFilter[];
        serviceRequest?: DeleteWhateversRequestFilter;
        serviceResponse?: DeleteWhateversResponseFilter;
        set?: boolean;
    }
    
    export interface PipeFeedsContextFilter {
        and?: PipeFeedsContextFilter[];
        delete?: PipeDeleteFeedsContextFilter;
        get?: PipeGetFeedsContextFilter;
        hash?: string;
        not?: PipeFeedsContextFilter[];
        or?: PipeFeedsContextFilter[];
        post?: PipePostFeedsContextFilter;
        put?: PipePutFeedsContextFilter;
        set?: boolean;
    }
    
    export interface PipeFeedsEndpoint {
        filter?: PipeFeedsRequestFilter;
        hash?: string;
    }
    
    export interface PipeFeedsEndpointFilter {
        and?: PipeFeedsEndpointFilter[];
        hash?: string;
        not?: PipeFeedsEndpointFilter[];
        or?: PipeFeedsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PipeFeedsRequestFilter {
        and?: PipeFeedsRequestFilter[];
        context?: PipeFeedsContextFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PipeModeFilter;
        not?: PipeFeedsRequestFilter[];
        or?: PipeFeedsRequestFilter[];
        set?: boolean;
    }
    
    export interface PipeGetAttachmentsContextFilter {
        and?: PipeGetAttachmentsContextFilter[];
        clientRequest?: GetAttachmentsRequestFilter;
        clientResponse?: GetAttachmentsResponseFilter;
        hash?: string;
        not?: PipeGetAttachmentsContextFilter[];
        or?: PipeGetAttachmentsContextFilter[];
        serviceRequest?: GetAttachmentsRequestFilter;
        serviceResponse?: GetAttachmentsResponseFilter;
        set?: boolean;
    }
    
    export interface PipeGetBlueWhateversContextFilter {
        and?: PipeGetBlueWhateversContextFilter[];
        clientRequest?: GetBlueWhateversRequestFilter;
        clientResponse?: GetBlueWhateversResponseFilter;
        hash?: string;
        not?: PipeGetBlueWhateversContextFilter[];
        or?: PipeGetBlueWhateversContextFilter[];
        serviceRequest?: GetBlueWhateversRequestFilter;
        serviceResponse?: GetBlueWhateversResponseFilter;
        set?: boolean;
    }
    
    export interface PipeGetClientAccountsContextFilter {
        and?: PipeGetClientAccountsContextFilter[];
        clientRequest?: GetClientAccountsRequestFilter;
        clientResponse?: GetClientAccountsResponseFilter;
        hash?: string;
        not?: PipeGetClientAccountsContextFilter[];
        or?: PipeGetClientAccountsContextFilter[];
        serviceRequest?: GetClientAccountsRequestFilter;
        serviceResponse?: GetClientAccountsResponseFilter;
        set?: boolean;
    }
    
    export interface PipeGetFeedsContextFilter {
        and?: PipeGetFeedsContextFilter[];
        clientRequest?: GetFeedsRequestFilter;
        clientResponse?: GetFeedsResponseFilter;
        hash?: string;
        not?: PipeGetFeedsContextFilter[];
        or?: PipeGetFeedsContextFilter[];
        serviceRequest?: GetFeedsRequestFilter;
        serviceResponse?: GetFeedsResponseFilter;
        set?: boolean;
    }
    
    export interface PipeGetPeopleContextFilter {
        and?: PipeGetPeopleContextFilter[];
        clientRequest?: GetPeopleRequestFilter;
        clientResponse?: GetPeopleResponseFilter;
        hash?: string;
        not?: PipeGetPeopleContextFilter[];
        or?: PipeGetPeopleContextFilter[];
        serviceRequest?: GetPeopleRequestFilter;
        serviceResponse?: GetPeopleResponseFilter;
        set?: boolean;
    }
    
    export interface PipeGetServiceAccountsContextFilter {
        and?: PipeGetServiceAccountsContextFilter[];
        clientRequest?: GetServiceAccountsRequestFilter;
        clientResponse?: GetServiceAccountsResponseFilter;
        hash?: string;
        not?: PipeGetServiceAccountsContextFilter[];
        or?: PipeGetServiceAccountsContextFilter[];
        serviceRequest?: GetServiceAccountsRequestFilter;
        serviceResponse?: GetServiceAccountsResponseFilter;
        set?: boolean;
    }
    
    export interface PipeGetServicesContextFilter {
        and?: PipeGetServicesContextFilter[];
        clientRequest?: GetServicesRequestFilter;
        clientResponse?: GetServicesResponseFilter;
        hash?: string;
        not?: PipeGetServicesContextFilter[];
        or?: PipeGetServicesContextFilter[];
        serviceRequest?: GetServicesRequestFilter;
        serviceResponse?: GetServicesResponseFilter;
        set?: boolean;
    }
    
    export interface PipeGetStatusesContextFilter {
        and?: PipeGetStatusesContextFilter[];
        clientRequest?: GetStatusesRequestFilter;
        clientResponse?: GetStatusesResponseFilter;
        hash?: string;
        not?: PipeGetStatusesContextFilter[];
        or?: PipeGetStatusesContextFilter[];
        serviceRequest?: GetStatusesRequestFilter;
        serviceResponse?: GetStatusesResponseFilter;
        set?: boolean;
    }
    
    export interface PipeGetWhateversContextFilter {
        and?: PipeGetWhateversContextFilter[];
        clientRequest?: GetWhateversRequestFilter;
        clientResponse?: GetWhateversResponseFilter;
        hash?: string;
        not?: PipeGetWhateversContextFilter[];
        or?: PipeGetWhateversContextFilter[];
        serviceRequest?: GetWhateversRequestFilter;
        serviceResponse?: GetWhateversResponseFilter;
        set?: boolean;
    }
    
    export interface PipeModeFilter {
        and?: PipeModeFilter[];
        context?: ContextPipeModeFilter;
        hash?: string;
        kind?: EnumFilter;
        not?: PipeModeFilter[];
        or?: PipeModeFilter[];
        set?: boolean;
    }
    
    export interface PipePeopleContextFilter {
        and?: PipePeopleContextFilter[];
        delete?: PipeDeletePeopleContextFilter;
        get?: PipeGetPeopleContextFilter;
        hash?: string;
        not?: PipePeopleContextFilter[];
        or?: PipePeopleContextFilter[];
        post?: PipePostPeopleContextFilter;
        put?: PipePutPeopleContextFilter;
        set?: boolean;
    }
    
    export interface PipePeopleEndpoint {
        filter?: PipePeopleRequestFilter;
        hash?: string;
    }
    
    export interface PipePeopleEndpointFilter {
        and?: PipePeopleEndpointFilter[];
        hash?: string;
        not?: PipePeopleEndpointFilter[];
        or?: PipePeopleEndpointFilter[];
        set?: boolean;
    }
    
    export interface PipePeopleRequestFilter {
        and?: PipePeopleRequestFilter[];
        context?: PipePeopleContextFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PipeModeFilter;
        not?: PipePeopleRequestFilter[];
        or?: PipePeopleRequestFilter[];
        set?: boolean;
    }
    
    export interface PipePostAttachmentsContextFilter {
        and?: PipePostAttachmentsContextFilter[];
        clientRequest?: PostAttachmentsRequestFilter;
        clientResponse?: PostAttachmentsResponseFilter;
        hash?: string;
        not?: PipePostAttachmentsContextFilter[];
        or?: PipePostAttachmentsContextFilter[];
        serviceRequest?: PostAttachmentsRequestFilter;
        serviceResponse?: PostAttachmentsResponseFilter;
        set?: boolean;
    }
    
    export interface PipePostBlueWhateversContextFilter {
        and?: PipePostBlueWhateversContextFilter[];
        clientRequest?: PostBlueWhateversRequestFilter;
        clientResponse?: PostBlueWhateversResponseFilter;
        hash?: string;
        not?: PipePostBlueWhateversContextFilter[];
        or?: PipePostBlueWhateversContextFilter[];
        serviceRequest?: PostBlueWhateversRequestFilter;
        serviceResponse?: PostBlueWhateversResponseFilter;
        set?: boolean;
    }
    
    export interface PipePostClientAccountsContextFilter {
        and?: PipePostClientAccountsContextFilter[];
        clientRequest?: PostClientAccountsRequestFilter;
        clientResponse?: PostClientAccountsResponseFilter;
        hash?: string;
        not?: PipePostClientAccountsContextFilter[];
        or?: PipePostClientAccountsContextFilter[];
        serviceRequest?: PostClientAccountsRequestFilter;
        serviceResponse?: PostClientAccountsResponseFilter;
        set?: boolean;
    }
    
    export interface PipePostFeedsContextFilter {
        and?: PipePostFeedsContextFilter[];
        clientRequest?: PostFeedsRequestFilter;
        clientResponse?: PostFeedsResponseFilter;
        hash?: string;
        not?: PipePostFeedsContextFilter[];
        or?: PipePostFeedsContextFilter[];
        serviceRequest?: PostFeedsRequestFilter;
        serviceResponse?: PostFeedsResponseFilter;
        set?: boolean;
    }
    
    export interface PipePostPeopleContextFilter {
        and?: PipePostPeopleContextFilter[];
        clientRequest?: PostPeopleRequestFilter;
        clientResponse?: PostPeopleResponseFilter;
        hash?: string;
        not?: PipePostPeopleContextFilter[];
        or?: PipePostPeopleContextFilter[];
        serviceRequest?: PostPeopleRequestFilter;
        serviceResponse?: PostPeopleResponseFilter;
        set?: boolean;
    }
    
    export interface PipePostServiceAccountsContextFilter {
        and?: PipePostServiceAccountsContextFilter[];
        clientRequest?: PostServiceAccountsRequestFilter;
        clientResponse?: PostServiceAccountsResponseFilter;
        hash?: string;
        not?: PipePostServiceAccountsContextFilter[];
        or?: PipePostServiceAccountsContextFilter[];
        serviceRequest?: PostServiceAccountsRequestFilter;
        serviceResponse?: PostServiceAccountsResponseFilter;
        set?: boolean;
    }
    
    export interface PipePostServicesContextFilter {
        and?: PipePostServicesContextFilter[];
        clientRequest?: PostServicesRequestFilter;
        clientResponse?: PostServicesResponseFilter;
        hash?: string;
        not?: PipePostServicesContextFilter[];
        or?: PipePostServicesContextFilter[];
        serviceRequest?: PostServicesRequestFilter;
        serviceResponse?: PostServicesResponseFilter;
        set?: boolean;
    }
    
    export interface PipePostStatusesContextFilter {
        and?: PipePostStatusesContextFilter[];
        clientRequest?: PostStatusesRequestFilter;
        clientResponse?: PostStatusesResponseFilter;
        hash?: string;
        not?: PipePostStatusesContextFilter[];
        or?: PipePostStatusesContextFilter[];
        serviceRequest?: PostStatusesRequestFilter;
        serviceResponse?: PostStatusesResponseFilter;
        set?: boolean;
    }
    
    export interface PipePostWhateversContextFilter {
        and?: PipePostWhateversContextFilter[];
        clientRequest?: PostWhateversRequestFilter;
        clientResponse?: PostWhateversResponseFilter;
        hash?: string;
        not?: PipePostWhateversContextFilter[];
        or?: PipePostWhateversContextFilter[];
        serviceRequest?: PostWhateversRequestFilter;
        serviceResponse?: PostWhateversResponseFilter;
        set?: boolean;
    }
    
    export interface PipePutAttachmentsContextFilter {
        and?: PipePutAttachmentsContextFilter[];
        clientRequest?: PutAttachmentsRequestFilter;
        clientResponse?: PutAttachmentsResponseFilter;
        hash?: string;
        not?: PipePutAttachmentsContextFilter[];
        or?: PipePutAttachmentsContextFilter[];
        serviceRequest?: PutAttachmentsRequestFilter;
        serviceResponse?: PutAttachmentsResponseFilter;
        set?: boolean;
    }
    
    export interface PipePutBlueWhateversContextFilter {
        and?: PipePutBlueWhateversContextFilter[];
        clientRequest?: PutBlueWhateversRequestFilter;
        clientResponse?: PutBlueWhateversResponseFilter;
        hash?: string;
        not?: PipePutBlueWhateversContextFilter[];
        or?: PipePutBlueWhateversContextFilter[];
        serviceRequest?: PutBlueWhateversRequestFilter;
        serviceResponse?: PutBlueWhateversResponseFilter;
        set?: boolean;
    }
    
    export interface PipePutClientAccountsContextFilter {
        and?: PipePutClientAccountsContextFilter[];
        clientRequest?: PutClientAccountsRequestFilter;
        clientResponse?: PutClientAccountsResponseFilter;
        hash?: string;
        not?: PipePutClientAccountsContextFilter[];
        or?: PipePutClientAccountsContextFilter[];
        serviceRequest?: PutClientAccountsRequestFilter;
        serviceResponse?: PutClientAccountsResponseFilter;
        set?: boolean;
    }
    
    export interface PipePutFeedsContextFilter {
        and?: PipePutFeedsContextFilter[];
        clientRequest?: PutFeedsRequestFilter;
        clientResponse?: PutFeedsResponseFilter;
        hash?: string;
        not?: PipePutFeedsContextFilter[];
        or?: PipePutFeedsContextFilter[];
        serviceRequest?: PutFeedsRequestFilter;
        serviceResponse?: PutFeedsResponseFilter;
        set?: boolean;
    }
    
    export interface PipePutPeopleContextFilter {
        and?: PipePutPeopleContextFilter[];
        clientRequest?: PutPeopleRequestFilter;
        clientResponse?: PutPeopleResponseFilter;
        hash?: string;
        not?: PipePutPeopleContextFilter[];
        or?: PipePutPeopleContextFilter[];
        serviceRequest?: PutPeopleRequestFilter;
        serviceResponse?: PutPeopleResponseFilter;
        set?: boolean;
    }
    
    export interface PipePutServiceAccountsContextFilter {
        and?: PipePutServiceAccountsContextFilter[];
        clientRequest?: PutServiceAccountsRequestFilter;
        clientResponse?: PutServiceAccountsResponseFilter;
        hash?: string;
        not?: PipePutServiceAccountsContextFilter[];
        or?: PipePutServiceAccountsContextFilter[];
        serviceRequest?: PutServiceAccountsRequestFilter;
        serviceResponse?: PutServiceAccountsResponseFilter;
        set?: boolean;
    }
    
    export interface PipePutServicesContextFilter {
        and?: PipePutServicesContextFilter[];
        clientRequest?: PutServicesRequestFilter;
        clientResponse?: PutServicesResponseFilter;
        hash?: string;
        not?: PipePutServicesContextFilter[];
        or?: PipePutServicesContextFilter[];
        serviceRequest?: PutServicesRequestFilter;
        serviceResponse?: PutServicesResponseFilter;
        set?: boolean;
    }
    
    export interface PipePutStatusesContextFilter {
        and?: PipePutStatusesContextFilter[];
        clientRequest?: PutStatusesRequestFilter;
        clientResponse?: PutStatusesResponseFilter;
        hash?: string;
        not?: PipePutStatusesContextFilter[];
        or?: PipePutStatusesContextFilter[];
        serviceRequest?: PutStatusesRequestFilter;
        serviceResponse?: PutStatusesResponseFilter;
        set?: boolean;
    }
    
    export interface PipePutWhateversContextFilter {
        and?: PipePutWhateversContextFilter[];
        clientRequest?: PutWhateversRequestFilter;
        clientResponse?: PutWhateversResponseFilter;
        hash?: string;
        not?: PipePutWhateversContextFilter[];
        or?: PipePutWhateversContextFilter[];
        serviceRequest?: PutWhateversRequestFilter;
        serviceResponse?: PutWhateversResponseFilter;
        set?: boolean;
    }
    
    export interface PipeServiceAccountsContextFilter {
        and?: PipeServiceAccountsContextFilter[];
        delete?: PipeDeleteServiceAccountsContextFilter;
        get?: PipeGetServiceAccountsContextFilter;
        hash?: string;
        not?: PipeServiceAccountsContextFilter[];
        or?: PipeServiceAccountsContextFilter[];
        post?: PipePostServiceAccountsContextFilter;
        put?: PipePutServiceAccountsContextFilter;
        set?: boolean;
    }
    
    export interface PipeServiceAccountsEndpoint {
        filter?: PipeServiceAccountsRequestFilter;
        hash?: string;
    }
    
    export interface PipeServiceAccountsEndpointFilter {
        and?: PipeServiceAccountsEndpointFilter[];
        hash?: string;
        not?: PipeServiceAccountsEndpointFilter[];
        or?: PipeServiceAccountsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PipeServiceAccountsRequestFilter {
        and?: PipeServiceAccountsRequestFilter[];
        context?: PipeServiceAccountsContextFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PipeModeFilter;
        not?: PipeServiceAccountsRequestFilter[];
        or?: PipeServiceAccountsRequestFilter[];
        set?: boolean;
    }
    
    export interface PipeServicesContextFilter {
        and?: PipeServicesContextFilter[];
        delete?: PipeDeleteServicesContextFilter;
        get?: PipeGetServicesContextFilter;
        hash?: string;
        not?: PipeServicesContextFilter[];
        or?: PipeServicesContextFilter[];
        post?: PipePostServicesContextFilter;
        put?: PipePutServicesContextFilter;
        set?: boolean;
    }
    
    export interface PipeServicesEndpoint {
        filter?: PipeServicesRequestFilter;
        hash?: string;
    }
    
    export interface PipeServicesEndpointFilter {
        and?: PipeServicesEndpointFilter[];
        hash?: string;
        not?: PipeServicesEndpointFilter[];
        or?: PipeServicesEndpointFilter[];
        set?: boolean;
    }
    
    export interface PipeServicesRequestFilter {
        and?: PipeServicesRequestFilter[];
        context?: PipeServicesContextFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PipeModeFilter;
        not?: PipeServicesRequestFilter[];
        or?: PipeServicesRequestFilter[];
        set?: boolean;
    }
    
    export interface PipeStatusesContextFilter {
        and?: PipeStatusesContextFilter[];
        delete?: PipeDeleteStatusesContextFilter;
        get?: PipeGetStatusesContextFilter;
        hash?: string;
        not?: PipeStatusesContextFilter[];
        or?: PipeStatusesContextFilter[];
        post?: PipePostStatusesContextFilter;
        put?: PipePutStatusesContextFilter;
        set?: boolean;
    }
    
    export interface PipeStatusesEndpoint {
        filter?: PipeStatusesRequestFilter;
        hash?: string;
    }
    
    export interface PipeStatusesEndpointFilter {
        and?: PipeStatusesEndpointFilter[];
        hash?: string;
        not?: PipeStatusesEndpointFilter[];
        or?: PipeStatusesEndpointFilter[];
        set?: boolean;
    }
    
    export interface PipeStatusesRequestFilter {
        and?: PipeStatusesRequestFilter[];
        context?: PipeStatusesContextFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PipeModeFilter;
        not?: PipeStatusesRequestFilter[];
        or?: PipeStatusesRequestFilter[];
        set?: boolean;
    }
    
    export interface PipeWhateversContextFilter {
        and?: PipeWhateversContextFilter[];
        delete?: PipeDeleteWhateversContextFilter;
        get?: PipeGetWhateversContextFilter;
        hash?: string;
        not?: PipeWhateversContextFilter[];
        or?: PipeWhateversContextFilter[];
        post?: PipePostWhateversContextFilter;
        put?: PipePutWhateversContextFilter;
        set?: boolean;
    }
    
    export interface PipeWhateversEndpoint {
        filter?: PipeWhateversRequestFilter;
        hash?: string;
    }
    
    export interface PipeWhateversEndpointFilter {
        and?: PipeWhateversEndpointFilter[];
        hash?: string;
        not?: PipeWhateversEndpointFilter[];
        or?: PipeWhateversEndpointFilter[];
        set?: boolean;
    }
    
    export interface PipeWhateversRequestFilter {
        and?: PipeWhateversRequestFilter[];
        context?: PipeWhateversContextFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PipeModeFilter;
        not?: PipeWhateversRequestFilter[];
        or?: PipeWhateversRequestFilter[];
        set?: boolean;
    }
    
    export interface PostAttachmentsEndpoint {
        filter?: PostAttachmentsRequestFilter;
        hash?: string;
    }
    
    export interface PostAttachmentsEndpointFilter {
        and?: PostAttachmentsEndpointFilter[];
        hash?: string;
        not?: PostAttachmentsEndpointFilter[];
        or?: PostAttachmentsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PostAttachmentsRequestFilter {
        and?: PostAttachmentsRequestFilter[];
        attachments?: AttachmentListFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PostModeFilter;
        not?: PostAttachmentsRequestFilter[];
        or?: PostAttachmentsRequestFilter[];
        set?: boolean;
    }
    
    export interface PostAttachmentsResponseFilter {
        and?: PostAttachmentsResponseFilter[];
        attachments?: AttachmentListFilter;
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PostAttachmentsResponseFilter[];
        or?: PostAttachmentsResponseFilter[];
        set?: boolean;
    }
    
    export interface PostBlueWhateversEndpoint {
        filter?: PostBlueWhateversRequestFilter;
        hash?: string;
    }
    
    export interface PostBlueWhateversEndpointFilter {
        and?: PostBlueWhateversEndpointFilter[];
        hash?: string;
        not?: PostBlueWhateversEndpointFilter[];
        or?: PostBlueWhateversEndpointFilter[];
        set?: boolean;
    }
    
    export interface PostBlueWhateversRequestFilter {
        and?: PostBlueWhateversRequestFilter[];
        blueWhatevers?: BlueWhateverListFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PostModeFilter;
        not?: PostBlueWhateversRequestFilter[];
        or?: PostBlueWhateversRequestFilter[];
        set?: boolean;
    }
    
    export interface PostBlueWhateversResponseFilter {
        and?: PostBlueWhateversResponseFilter[];
        blueWhatevers?: BlueWhateverListFilter;
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PostBlueWhateversResponseFilter[];
        or?: PostBlueWhateversResponseFilter[];
        set?: boolean;
    }
    
    export interface PostClientAccountsEndpoint {
        filter?: PostClientAccountsRequestFilter;
        hash?: string;
    }
    
    export interface PostClientAccountsEndpointFilter {
        and?: PostClientAccountsEndpointFilter[];
        hash?: string;
        not?: PostClientAccountsEndpointFilter[];
        or?: PostClientAccountsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PostClientAccountsRequest {
        auth?: Auth;
        clientAccounts?: ClientAccount[];
        hash?: string;
        meta?: RequestMeta;
        mode?: PostMode;
        select?: PostClientAccountsResponseSelect;
        serviceFilter?: ServiceFilter;
    }
    
    export interface PostClientAccountsRequestFilter {
        and?: PostClientAccountsRequestFilter[];
        clientAccounts?: ClientAccountListFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PostModeFilter;
        not?: PostClientAccountsRequestFilter[];
        or?: PostClientAccountsRequestFilter[];
        set?: boolean;
    }
    
    export interface PostClientAccountsResponse {
        clientAccounts?: ClientAccount[];
        hash?: string;
        meta?: ResponseMeta;
    }
    
    export interface PostClientAccountsResponseFilter {
        and?: PostClientAccountsResponseFilter[];
        clientAccounts?: ClientAccountListFilter;
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PostClientAccountsResponseFilter[];
        or?: PostClientAccountsResponseFilter[];
        set?: boolean;
    }
    
    export interface PostClientAccountsResponseSelect {
        clientAccounts?: ClientAccountSelect;
        hash?: string;
        meta?: ResponseMetaSelect;
        selectAll?: boolean;
    }
    
    export interface PostFeedsEndpoint {
        filter?: PostFeedsRequestFilter;
        hash?: string;
    }
    
    export interface PostFeedsEndpointFilter {
        and?: PostFeedsEndpointFilter[];
        hash?: string;
        not?: PostFeedsEndpointFilter[];
        or?: PostFeedsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PostFeedsRequestFilter {
        and?: PostFeedsRequestFilter[];
        feeds?: FeedListFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PostModeFilter;
        not?: PostFeedsRequestFilter[];
        or?: PostFeedsRequestFilter[];
        set?: boolean;
    }
    
    export interface PostFeedsResponseFilter {
        and?: PostFeedsResponseFilter[];
        feeds?: FeedListFilter;
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PostFeedsResponseFilter[];
        or?: PostFeedsResponseFilter[];
        set?: boolean;
    }
    
    export interface PostMode {
        collection?: CollectionPostMode;
        hash?: string;
        kind?: string;
    }
    
    export interface PostModeFilter {
        and?: PostModeFilter[];
        hash?: string;
        kind?: EnumFilter;
        not?: PostModeFilter[];
        or?: PostModeFilter[];
        set?: boolean;
    }
    
    export interface PostPeopleEndpoint {
        filter?: PostPeopleRequestFilter;
        hash?: string;
    }
    
    export interface PostPeopleEndpointFilter {
        and?: PostPeopleEndpointFilter[];
        hash?: string;
        not?: PostPeopleEndpointFilter[];
        or?: PostPeopleEndpointFilter[];
        set?: boolean;
    }
    
    export interface PostPeopleRequestFilter {
        and?: PostPeopleRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PostModeFilter;
        not?: PostPeopleRequestFilter[];
        or?: PostPeopleRequestFilter[];
        people?: PersonListFilter;
        set?: boolean;
    }
    
    export interface PostPeopleResponseFilter {
        and?: PostPeopleResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PostPeopleResponseFilter[];
        or?: PostPeopleResponseFilter[];
        people?: PersonListFilter;
        set?: boolean;
    }
    
    export interface PostServiceAccountsEndpoint {
        filter?: PostServiceAccountsRequestFilter;
        hash?: string;
    }
    
    export interface PostServiceAccountsEndpointFilter {
        and?: PostServiceAccountsEndpointFilter[];
        hash?: string;
        not?: PostServiceAccountsEndpointFilter[];
        or?: PostServiceAccountsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PostServiceAccountsRequest {
        auth?: Auth;
        hash?: string;
        meta?: RequestMeta;
        mode?: PostMode;
        select?: PostServiceAccountsResponseSelect;
        serviceAccounts?: ServiceAccount[];
        serviceFilter?: ServiceFilter;
    }
    
    export interface PostServiceAccountsRequestFilter {
        and?: PostServiceAccountsRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PostModeFilter;
        not?: PostServiceAccountsRequestFilter[];
        or?: PostServiceAccountsRequestFilter[];
        serviceAccounts?: ServiceAccountListFilter;
        set?: boolean;
    }
    
    export interface PostServiceAccountsResponse {
        hash?: string;
        meta?: ResponseMeta;
        serviceAccounts?: ServiceAccount[];
    }
    
    export interface PostServiceAccountsResponseFilter {
        and?: PostServiceAccountsResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PostServiceAccountsResponseFilter[];
        or?: PostServiceAccountsResponseFilter[];
        serviceAccounts?: ServiceAccountListFilter;
        set?: boolean;
    }
    
    export interface PostServiceAccountsResponseSelect {
        hash?: string;
        meta?: ResponseMetaSelect;
        selectAll?: boolean;
        serviceAccounts?: ServiceAccountSelect;
    }
    
    export interface PostServicesEndpoint {
        filter?: PostServicesRequestFilter;
        hash?: string;
    }
    
    export interface PostServicesEndpointFilter {
        and?: PostServicesEndpointFilter[];
        hash?: string;
        not?: PostServicesEndpointFilter[];
        or?: PostServicesEndpointFilter[];
        set?: boolean;
    }
    
    export interface PostServicesRequestFilter {
        and?: PostServicesRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PostModeFilter;
        not?: PostServicesRequestFilter[];
        or?: PostServicesRequestFilter[];
        services?: ServiceListFilter;
        set?: boolean;
    }
    
    export interface PostServicesResponseFilter {
        and?: PostServicesResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PostServicesResponseFilter[];
        or?: PostServicesResponseFilter[];
        services?: ServiceListFilter;
        set?: boolean;
    }
    
    export interface PostStatusesEndpoint {
        filter?: PostStatusesRequestFilter;
        hash?: string;
    }
    
    export interface PostStatusesEndpointFilter {
        and?: PostStatusesEndpointFilter[];
        hash?: string;
        not?: PostStatusesEndpointFilter[];
        or?: PostStatusesEndpointFilter[];
        set?: boolean;
    }
    
    export interface PostStatusesRequestFilter {
        and?: PostStatusesRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PostModeFilter;
        not?: PostStatusesRequestFilter[];
        or?: PostStatusesRequestFilter[];
        set?: boolean;
        statuses?: StatusListFilter;
    }
    
    export interface PostStatusesResponseFilter {
        and?: PostStatusesResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PostStatusesResponseFilter[];
        or?: PostStatusesResponseFilter[];
        set?: boolean;
        statuses?: StatusListFilter;
    }
    
    export interface PostWhateversEndpoint {
        filter?: PostWhateversRequestFilter;
        hash?: string;
    }
    
    export interface PostWhateversEndpointFilter {
        and?: PostWhateversEndpointFilter[];
        hash?: string;
        not?: PostWhateversEndpointFilter[];
        or?: PostWhateversEndpointFilter[];
        set?: boolean;
    }
    
    export interface PostWhateversRequestFilter {
        and?: PostWhateversRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PostModeFilter;
        not?: PostWhateversRequestFilter[];
        or?: PostWhateversRequestFilter[];
        set?: boolean;
        whatevers?: WhateverListFilter;
    }
    
    export interface PostWhateversResponseFilter {
        and?: PostWhateversResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PostWhateversResponseFilter[];
        or?: PostWhateversResponseFilter[];
        set?: boolean;
        whatevers?: WhateverListFilter;
    }
    
    export interface PutAttachmentsEndpoint {
        filter?: PutAttachmentsRequestFilter;
        hash?: string;
    }
    
    export interface PutAttachmentsEndpointFilter {
        and?: PutAttachmentsEndpointFilter[];
        hash?: string;
        not?: PutAttachmentsEndpointFilter[];
        or?: PutAttachmentsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PutAttachmentsRequestFilter {
        and?: PutAttachmentsRequestFilter[];
        attachments?: AttachmentListFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PutModeFilter;
        not?: PutAttachmentsRequestFilter[];
        or?: PutAttachmentsRequestFilter[];
        set?: boolean;
    }
    
    export interface PutAttachmentsResponseFilter {
        and?: PutAttachmentsResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PutAttachmentsResponseFilter[];
        or?: PutAttachmentsResponseFilter[];
        set?: boolean;
    }
    
    export interface PutBlueWhateversEndpoint {
        filter?: PutBlueWhateversRequestFilter;
        hash?: string;
    }
    
    export interface PutBlueWhateversEndpointFilter {
        and?: PutBlueWhateversEndpointFilter[];
        hash?: string;
        not?: PutBlueWhateversEndpointFilter[];
        or?: PutBlueWhateversEndpointFilter[];
        set?: boolean;
    }
    
    export interface PutBlueWhateversRequestFilter {
        and?: PutBlueWhateversRequestFilter[];
        blueWhatevers?: BlueWhateverListFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PutModeFilter;
        not?: PutBlueWhateversRequestFilter[];
        or?: PutBlueWhateversRequestFilter[];
        set?: boolean;
    }
    
    export interface PutBlueWhateversResponseFilter {
        and?: PutBlueWhateversResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PutBlueWhateversResponseFilter[];
        or?: PutBlueWhateversResponseFilter[];
        set?: boolean;
    }
    
    export interface PutClientAccountsEndpoint {
        filter?: PutClientAccountsRequestFilter;
        hash?: string;
    }
    
    export interface PutClientAccountsEndpointFilter {
        and?: PutClientAccountsEndpointFilter[];
        hash?: string;
        not?: PutClientAccountsEndpointFilter[];
        or?: PutClientAccountsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PutClientAccountsRequestFilter {
        and?: PutClientAccountsRequestFilter[];
        clientAccounts?: ClientAccountListFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PutModeFilter;
        not?: PutClientAccountsRequestFilter[];
        or?: PutClientAccountsRequestFilter[];
        set?: boolean;
    }
    
    export interface PutClientAccountsResponseFilter {
        and?: PutClientAccountsResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PutClientAccountsResponseFilter[];
        or?: PutClientAccountsResponseFilter[];
        set?: boolean;
    }
    
    export interface PutFeedsEndpoint {
        filter?: PutFeedsRequestFilter;
        hash?: string;
    }
    
    export interface PutFeedsEndpointFilter {
        and?: PutFeedsEndpointFilter[];
        hash?: string;
        not?: PutFeedsEndpointFilter[];
        or?: PutFeedsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PutFeedsRequestFilter {
        and?: PutFeedsRequestFilter[];
        feeds?: FeedListFilter;
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PutModeFilter;
        not?: PutFeedsRequestFilter[];
        or?: PutFeedsRequestFilter[];
        set?: boolean;
    }
    
    export interface PutFeedsResponseFilter {
        and?: PutFeedsResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PutFeedsResponseFilter[];
        or?: PutFeedsResponseFilter[];
        set?: boolean;
    }
    
    export interface PutModeFilter {
        and?: PutModeFilter[];
        hash?: string;
        kind?: EnumFilter;
        not?: PutModeFilter[];
        or?: PutModeFilter[];
        relation?: RelationPutModeFilter;
        set?: boolean;
    }
    
    export interface PutPeopleEndpoint {
        filter?: PutPeopleRequestFilter;
        hash?: string;
    }
    
    export interface PutPeopleEndpointFilter {
        and?: PutPeopleEndpointFilter[];
        hash?: string;
        not?: PutPeopleEndpointFilter[];
        or?: PutPeopleEndpointFilter[];
        set?: boolean;
    }
    
    export interface PutPeopleRequestFilter {
        and?: PutPeopleRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PutModeFilter;
        not?: PutPeopleRequestFilter[];
        or?: PutPeopleRequestFilter[];
        people?: PersonListFilter;
        set?: boolean;
    }
    
    export interface PutPeopleResponseFilter {
        and?: PutPeopleResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PutPeopleResponseFilter[];
        or?: PutPeopleResponseFilter[];
        set?: boolean;
    }
    
    export interface PutServiceAccountsEndpoint {
        filter?: PutServiceAccountsRequestFilter;
        hash?: string;
    }
    
    export interface PutServiceAccountsEndpointFilter {
        and?: PutServiceAccountsEndpointFilter[];
        hash?: string;
        not?: PutServiceAccountsEndpointFilter[];
        or?: PutServiceAccountsEndpointFilter[];
        set?: boolean;
    }
    
    export interface PutServiceAccountsRequestFilter {
        and?: PutServiceAccountsRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PutModeFilter;
        not?: PutServiceAccountsRequestFilter[];
        or?: PutServiceAccountsRequestFilter[];
        serviceAccounts?: ServiceAccountListFilter;
        set?: boolean;
    }
    
    export interface PutServiceAccountsResponseFilter {
        and?: PutServiceAccountsResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PutServiceAccountsResponseFilter[];
        or?: PutServiceAccountsResponseFilter[];
        set?: boolean;
    }
    
    export interface PutServicesEndpoint {
        filter?: PutServicesRequestFilter;
        hash?: string;
    }
    
    export interface PutServicesEndpointFilter {
        and?: PutServicesEndpointFilter[];
        hash?: string;
        not?: PutServicesEndpointFilter[];
        or?: PutServicesEndpointFilter[];
        set?: boolean;
    }
    
    export interface PutServicesRequestFilter {
        and?: PutServicesRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PutModeFilter;
        not?: PutServicesRequestFilter[];
        or?: PutServicesRequestFilter[];
        services?: ServiceListFilter;
        set?: boolean;
    }
    
    export interface PutServicesResponseFilter {
        and?: PutServicesResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PutServicesResponseFilter[];
        or?: PutServicesResponseFilter[];
        set?: boolean;
    }
    
    export interface PutStatusesEndpoint {
        filter?: PutStatusesRequestFilter;
        hash?: string;
    }
    
    export interface PutStatusesEndpointFilter {
        and?: PutStatusesEndpointFilter[];
        hash?: string;
        not?: PutStatusesEndpointFilter[];
        or?: PutStatusesEndpointFilter[];
        set?: boolean;
    }
    
    export interface PutStatusesRequestFilter {
        and?: PutStatusesRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PutModeFilter;
        not?: PutStatusesRequestFilter[];
        or?: PutStatusesRequestFilter[];
        set?: boolean;
        statuses?: StatusListFilter;
    }
    
    export interface PutStatusesResponseFilter {
        and?: PutStatusesResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PutStatusesResponseFilter[];
        or?: PutStatusesResponseFilter[];
        set?: boolean;
    }
    
    export interface PutWhateversEndpoint {
        filter?: PutWhateversRequestFilter;
        hash?: string;
    }
    
    export interface PutWhateversEndpointFilter {
        and?: PutWhateversEndpointFilter[];
        hash?: string;
        not?: PutWhateversEndpointFilter[];
        or?: PutWhateversEndpointFilter[];
        set?: boolean;
    }
    
    export interface PutWhateversRequestFilter {
        and?: PutWhateversRequestFilter[];
        hash?: string;
        meta?: RequestMetaFilter;
        mode?: PutModeFilter;
        not?: PutWhateversRequestFilter[];
        or?: PutWhateversRequestFilter[];
        set?: boolean;
        whatevers?: WhateverListFilter;
    }
    
    export interface PutWhateversResponseFilter {
        and?: PutWhateversResponseFilter[];
        hash?: string;
        meta?: ResponseMetaFilter;
        not?: PutWhateversResponseFilter[];
        or?: PutWhateversResponseFilter[];
        set?: boolean;
    }
    
    export interface RelationGetMode {
        hash?: string;
        id?: ServiceId;
        relation?: string;
    }
    
    export interface RelationGetModeFilter {
        and?: RelationGetModeFilter[];
        hash?: string;
        id?: ServiceIdFilter;
        not?: RelationGetModeFilter[];
        or?: RelationGetModeFilter[];
        relation?: StringFilter;
        set?: boolean;
    }
    
    export interface RelationPutModeFilter {
        and?: RelationPutModeFilter[];
        hash?: string;
        id?: ServiceIdFilter;
        ids?: ServiceIdListFilter;
        not?: RelationPutModeFilter[];
        operation?: EnumFilter;
        or?: RelationPutModeFilter[];
        relation?: StringFilter;
        set?: boolean;
    }
    
    export interface RequestMeta {
        createdAt?: Timestamp;
        hash?: string;
    }
    
    export interface RequestMetaFilter {
        and?: RequestMetaFilter[];
        createdAt?: TimestampFilter;
        hash?: string;
        not?: RequestMetaFilter[];
        or?: RequestMetaFilter[];
        set?: boolean;
    }
    
    export interface ResponseMeta {
        errors?: Error[];
        hash?: string;
        kind?: string;
        services?: Service[];
    }
    
    export interface ResponseMetaFilter {
        and?: ResponseMetaFilter[];
        errors?: ErrorListFilter;
        hash?: string;
        kind?: EnumFilter;
        not?: ResponseMetaFilter[];
        or?: ResponseMetaFilter[];
        services?: ServiceListFilter;
        set?: boolean;
    }
    
    export interface ResponseMetaSelect {
        errors?: ErrorSelect;
        hash?: string;
        kind?: boolean;
        selectAll?: boolean;
        services?: ServiceSelect;
    }
    
    export interface SearchGetMode {
        hash?: string;
        location?: LocationQuery;
        term?: string;
    }
    
    export interface SearchGetModeFilter {
        and?: SearchGetModeFilter[];
        hash?: string;
        location?: LocationQueryFilter;
        not?: SearchGetModeFilter[];
        or?: SearchGetModeFilter[];
        set?: boolean;
        term?: StringFilter;
    }
    
    export interface Service {
        alternativeIds?: Id[];
        endpoints?: Endpoints;
        hash?: string;
        id?: ServiceId;
        isVirtual?: boolean;
        meta?: TypeMeta;
        name?: string;
        port?: number;
        relations?: ServiceRelations;
        transport?: string;
        url?: Url;
    }
    
    export interface ServiceAccount {
        alternativeIds?: Id[];
        handle?: string;
        hash?: string;
        id?: ServiceId;
        meta?: TypeMeta;
        password?: Password;
        relations?: ServiceAccountRelations;
        url?: Url;
    }
    
    export interface ServiceAccountFilter {
        alternativeIds?: IdListFilter;
        and?: ServiceAccountFilter[];
        handle?: StringFilter;
        hash?: string;
        id?: ServiceIdFilter;
        meta?: TypeMetaFilter;
        not?: ServiceAccountFilter[];
        or?: ServiceAccountFilter[];
        password?: PasswordFilter;
        set?: boolean;
        url?: UrlFilter;
    }
    
    export interface ServiceAccountListFilter {
        every?: ServiceAccountFilter;
        hash?: string;
        none?: ServiceAccountFilter;
        some?: ServiceAccountFilter;
    }
    
    export interface ServiceAccountRelations {
        hash?: string;
        ownedByClientAccounts?: ClientAccountsCollection;
        usedByServices?: ServicesCollection;
    }
    
    export interface ServiceAccountRelationsSelect {
        hash?: string;
        ownedByClientAccounts?: ClientAccountsCollectionSelect;
        selectAll?: boolean;
        usedByServices?: ServicesCollectionSelect;
    }
    
    export interface ServiceAccountsCollection {
        hash?: string;
        meta?: CollectionMeta;
        serviceAccounts?: ServiceAccount[];
    }
    
    export interface ServiceAccountsCollectionSelect {
        hash?: string;
        meta?: CollectionMetaSelect;
        selectAll?: boolean;
        serviceAccounts?: ServiceAccountSelect;
    }
    
    export interface ServiceAccountSelect {
        alternativeIds?: IdSelect;
        handle?: boolean;
        hash?: string;
        id?: ServiceIdSelect;
        meta?: TypeMetaSelect;
        password?: PasswordSelect;
        relations?: ServiceAccountRelationsSelect;
        selectAll?: boolean;
        url?: UrlSelect;
    }
    
    export interface ServiceAccountSort {
        handle?: string;
        hash?: string;
        id?: ServiceIdSort;
        meta?: TypeMetaSort;
        password?: PasswordSort;
        url?: UrlSort;
    }
    
    export interface ServiceFilter {
        alternativeIds?: IdListFilter;
        and?: ServiceFilter[];
        endpoints?: EndpointsFilter;
        hash?: string;
        id?: ServiceIdFilter;
        isVirtual?: BoolFilter;
        meta?: TypeMetaFilter;
        name?: StringFilter;
        not?: ServiceFilter[];
        or?: ServiceFilter[];
        port?: Int32Filter;
        set?: boolean;
        transport?: EnumFilter;
        url?: UrlFilter;
    }
    
    export interface ServiceId {
        hash?: string;
        serviceName?: string;
        value?: string;
    }
    
    export interface ServiceIdFilter {
        and?: ServiceIdFilter[];
        hash?: string;
        not?: ServiceIdFilter[];
        or?: ServiceIdFilter[];
        serviceName?: StringFilter;
        set?: boolean;
        value?: StringFilter;
    }
    
    export interface ServiceIdListFilter {
        every?: ServiceIdFilter;
        hash?: string;
        none?: ServiceIdFilter;
        some?: ServiceIdFilter;
    }
    
    export interface ServiceIdSelect {
        hash?: string;
        selectAll?: boolean;
        serviceName?: boolean;
        value?: boolean;
    }
    
    export interface ServiceIdSort {
        hash?: string;
        serviceName?: string;
        value?: string;
    }
    
    export interface ServiceListFilter {
        every?: ServiceFilter;
        hash?: string;
        none?: ServiceFilter;
        some?: ServiceFilter;
    }
    
    export interface ServicePage {
        hash?: string;
        page?: Page;
        service?: Service;
    }
    
    export interface ServicePageFilter {
        and?: ServicePageFilter[];
        hash?: string;
        not?: ServicePageFilter[];
        or?: ServicePageFilter[];
        page?: PageFilter;
        service?: ServiceFilter;
        set?: boolean;
    }
    
    export interface ServicePageListFilter {
        every?: ServicePageFilter;
        hash?: string;
        none?: ServicePageFilter;
        some?: ServicePageFilter;
    }
    
    export interface ServicePageSelect {
        hash?: string;
        page?: PageSelect;
        selectAll?: boolean;
        service?: ServiceSelect;
    }
    
    export interface ServiceRelations {
        hash?: string;
        usesServiceAccounts?: ServiceAccountsCollection;
    }
    
    export interface ServiceRelationsSelect {
        hash?: string;
        selectAll?: boolean;
        usesServiceAccounts?: ServiceAccountsCollectionSelect;
    }
    
    export interface ServicesCollection {
        hash?: string;
        meta?: CollectionMeta;
        services?: Service[];
    }
    
    export interface ServicesCollectionSelect {
        hash?: string;
        meta?: CollectionMetaSelect;
        selectAll?: boolean;
        services?: ServiceSelect;
    }
    
    export interface ServiceSelect {
        alternativeIds?: IdSelect;
        hash?: string;
        id?: ServiceIdSelect;
        isVirtual?: boolean;
        meta?: TypeMetaSelect;
        name?: boolean;
        port?: boolean;
        relations?: ServiceRelationsSelect;
        selectAll?: boolean;
        transport?: boolean;
        url?: UrlSelect;
    }
    
    export interface ServiceSort {
        hash?: string;
        id?: ServiceIdSort;
        isVirtual?: string;
        meta?: TypeMetaSort;
        name?: string;
        port?: string;
        url?: UrlSort;
    }
    
    export interface StatusFilter {
        alternativeIds?: IdListFilter;
        and?: StatusFilter[];
        content?: TextFilter;
        hash?: string;
        id?: ServiceIdFilter;
        meta?: TypeMetaFilter;
        not?: StatusFilter[];
        or?: StatusFilter[];
        pinned?: BoolFilter;
        sensitive?: BoolFilter;
        set?: boolean;
        spoilerText?: TextFilter;
    }
    
    export interface StatusListFilter {
        every?: StatusFilter;
        hash?: string;
        none?: StatusFilter;
        some?: StatusFilter;
    }
    
    export interface StringFilter {
        and?: StringFilter[];
        caseSensitive?: boolean;
        contains?: string;
        endsWith?: string;
        hash?: string;
        in?: string[];
        is?: string;
        not?: string;
        notContains?: string;
        notEndsWith?: string;
        notIn?: string[];
        notStartsWith?: string;
        or?: StringFilter[];
        set?: boolean;
        startsWith?: string;
    }
    
    export interface StringListFilter {
        and?: StringFilter;
        hash?: string;
        not?: StringFilter;
        or?: StringFilter;
    }
    
    export interface Text {
        formatting?: string;
        hash?: string;
        language?: string;
        value?: string;
    }
    
    export interface TextFilter {
        and?: TextFilter[];
        formatting?: EnumFilter;
        hash?: string;
        language?: EnumFilter;
        not?: TextFilter[];
        or?: TextFilter[];
        set?: boolean;
        value?: StringFilter;
    }
    
    export interface TextSelect {
        formatting?: boolean;
        hash?: string;
        language?: boolean;
        selectAll?: boolean;
        value?: boolean;
    }
    
    export interface Timestamp {
        hash?: string;
        kind?: string;
        value?: string;
    }
    
    export interface TimestampFilter {
        and?: TimestampFilter[];
        hash?: string;
        kind?: EnumFilter;
        not?: TimestampFilter[];
        or?: TimestampFilter[];
        set?: boolean;
        value?: StringFilter;
    }
    
    export interface TimestampSelect {
        hash?: string;
        kind?: boolean;
        selectAll?: boolean;
        value?: boolean;
    }
    
    export interface TimestampSort {
        hash?: string;
        value?: string;
    }
    
    export interface Token {
        hash?: string;
        value?: string;
    }
    
    export interface TokenFilter {
        and?: TokenFilter[];
        hash?: string;
        not?: TokenFilter[];
        or?: TokenFilter[];
        set?: boolean;
        value?: StringFilter;
    }
    
    export interface TokenSelect {
        hash?: string;
        selectAll?: boolean;
        value?: boolean;
    }
    
    export interface TypeMeta {
        archived?: boolean;
        createdAt?: Timestamp;
        deletedAt?: Timestamp;
        hash?: string;
        sensitive?: boolean;
        service?: Service;
        updateAt?: Timestamp;
    }
    
    export interface TypeMetaFilter {
        and?: TypeMetaFilter[];
        archived?: BoolFilter;
        createdAt?: TimestampFilter;
        deletedAt?: TimestampFilter;
        hash?: string;
        not?: TypeMetaFilter[];
        or?: TypeMetaFilter[];
        sensitive?: BoolFilter;
        service?: ServiceFilter;
        set?: boolean;
        updateAt?: TimestampFilter;
    }
    
    export interface TypeMetaSelect {
        archived?: boolean;
        createdAt?: TimestampSelect;
        deletedAt?: TimestampSelect;
        hash?: string;
        selectAll?: boolean;
        sensitive?: boolean;
        service?: ServiceSelect;
        updateAt?: TimestampSelect;
    }
    
    export interface TypeMetaSort {
        archived?: string;
        createdAt?: TimestampSort;
        deletedAt?: TimestampSort;
        hash?: string;
        sensitive?: string;
        service?: ServiceSort;
        updateAt?: TimestampSort;
    }
    
    export interface Url {
        hash?: string;
        value?: string;
    }
    
    export interface UrlFilter {
        and?: UrlFilter[];
        hash?: string;
        not?: UrlFilter[];
        or?: UrlFilter[];
        set?: boolean;
        value?: StringFilter;
    }
    
    export interface UrlSelect {
        hash?: string;
        selectAll?: boolean;
        value?: boolean;
    }
    
    export interface UrlSort {
        hash?: string;
        value?: string;
    }
    
    export interface VerifyTokenEndpoint {
        filter?: VerifyTokenRequestFilter;
        hash?: string;
    }
    
    export interface VerifyTokenEndpointFilter {
        and?: VerifyTokenEndpointFilter[];
        hash?: string;
        not?: VerifyTokenEndpointFilter[];
        or?: VerifyTokenEndpointFilter[];
        set?: boolean;
    }
    
    export interface VerifyTokenInputFilter {
        and?: VerifyTokenInputFilter[];
        hash?: string;
        not?: VerifyTokenInputFilter[];
        or?: VerifyTokenInputFilter[];
        set?: boolean;
        token?: TokenFilter;
    }
    
    export interface VerifyTokenRequestFilter {
        and?: VerifyTokenRequestFilter[];
        hash?: string;
        input?: VerifyTokenInputFilter;
        meta?: RequestMetaFilter;
        not?: VerifyTokenRequestFilter[];
        or?: VerifyTokenRequestFilter[];
        set?: boolean;
    }
    
    export interface Whatever {
        alternativeIds?: Id[];
        boolField?: boolean;
        boolList?: boolean[];
        enumField?: string;
        enumList?: string[];
        float64Field?: number;
        float64List?: number[];
        hash?: string;
        id?: ServiceId;
        int32Field?: number;
        int32List?: number[];
        meta?: TypeMeta;
        relations?: WhateverRelations;
        stringField?: string;
        stringList?: string[];
        unionField?: WhateverUnion;
        unionList?: WhateverUnion[];
    }
    
    export interface WhateverFilter {
        alternativeIds?: IdListFilter;
        and?: WhateverFilter[];
        boolField?: BoolFilter;
        boolList?: BoolListFilter;
        enumField?: EnumFilter;
        enumList?: EnumListFilter;
        float64Field?: Float64Filter;
        float64List?: Float64ListFilter;
        hash?: string;
        id?: ServiceIdFilter;
        int32Field?: Int32Filter;
        int32List?: Int32ListFilter;
        meta?: TypeMetaFilter;
        not?: WhateverFilter[];
        or?: WhateverFilter[];
        set?: boolean;
        stringField?: StringFilter;
        stringList?: StringListFilter;
        unionField?: WhateverUnionFilter;
        unionList?: WhateverUnionListFilter;
    }
    
    export interface WhateverListFilter {
        every?: WhateverFilter;
        hash?: string;
        none?: WhateverFilter;
        some?: WhateverFilter;
    }
    
    export interface WhateverRelations {
        hash?: string;
        knewByWhatevers?: WhateversCollection;
        knowsBlueWhatevers?: BlueWhateversCollection;
        knowsWhatevers?: WhateversCollection;
    }
    
    export interface WhateverRelationsSelect {
        hash?: string;
        knewByWhatevers?: WhateversCollectionSelect;
        knowsBlueWhatevers?: BlueWhateversCollectionSelect;
        knowsWhatevers?: WhateversCollectionSelect;
        selectAll?: boolean;
    }
    
    export interface WhateversCollection {
        hash?: string;
        meta?: CollectionMeta;
        whatevers?: Whatever[];
    }
    
    export interface WhateversCollectionSelect {
        hash?: string;
        meta?: CollectionMetaSelect;
        selectAll?: boolean;
        whatevers?: WhateverSelect;
    }
    
    export interface WhateverSelect {
        alternativeIds?: IdSelect;
        boolField?: boolean;
        boolList?: boolean;
        enumField?: boolean;
        enumList?: boolean;
        float64Field?: boolean;
        float64List?: boolean;
        hash?: string;
        id?: ServiceIdSelect;
        int32Field?: boolean;
        int32List?: boolean;
        meta?: TypeMetaSelect;
        relations?: WhateverRelationsSelect;
        selectAll?: boolean;
        stringField?: boolean;
        stringList?: boolean;
        unionField?: WhateverUnionSelect;
        unionList?: WhateverUnionSelect;
    }
    
    export interface WhateverSort {
        boolField?: string;
        float64Field?: string;
        hash?: string;
        id?: ServiceIdSort;
        int32Field?: string;
        meta?: TypeMetaSort;
        stringField?: string;
        unionField?: WhateverUnionSort;
    }
    
    export interface WhateverUnion {
        boolField?: boolean;
        enumField?: string;
        float64Field?: number;
        hash?: string;
        int32Field?: number;
        kind?: string;
        stringField?: string;
    }
    
    export interface WhateverUnionFilter {
        and?: WhateverUnionFilter[];
        boolField?: BoolFilter;
        enumField?: EnumFilter;
        float64Field?: Float64Filter;
        hash?: string;
        int32Field?: Int32Filter;
        kind?: EnumFilter;
        not?: WhateverUnionFilter[];
        or?: WhateverUnionFilter[];
        set?: boolean;
        stringField?: StringFilter;
    }
    
    export interface WhateverUnionListFilter {
        every?: WhateverUnionFilter;
        hash?: string;
        none?: WhateverUnionFilter;
        some?: WhateverUnionFilter;
    }
    
    export interface WhateverUnionSelect {
        boolField?: boolean;
        enumField?: boolean;
        float64Field?: boolean;
        hash?: string;
        int32Field?: boolean;
        kind?: boolean;
        selectAll?: boolean;
        stringField?: boolean;
    }
    
    export interface WhateverUnionSort {
        boolField?: string;
        float64Field?: string;
        hash?: string;
        int32Field?: string;
        stringField?: string;
    }
    
}