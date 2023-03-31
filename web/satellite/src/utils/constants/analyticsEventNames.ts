// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

// Make sure these event names match up with the client-side event names in satellite/analytics/service.go
export enum AnalyticsEvent {
    GATEWAY_CREDENTIALS_CREATED = 'Credentials Created',
    PASSPHRASE_CREATED = 'Passphrase Created',
    EXTERNAL_LINK_CLICKED = 'External Link Clicked',
    PATH_SELECTED = 'Path Selected',
    LINK_SHARED = 'Link Shared',
    OBJECT_UPLOADED = 'Object Uploaded',
    API_KEY_GENERATED = 'API Key Generated',
    UPGRADE_BANNER_CLICKED = 'Upgrade Banner Clicked',
    MODAL_ADD_CARD = 'Credit Card Added In Modal',
    MODAL_ADD_TOKENS = 'Storj Token Added In Modal',
    SEARCH_BUCKETS = 'Search Buckets',
    NAVIGATE_PROJECTS = 'Navigate Projects',
    MANAGE_PROJECTS_CLICKED = 'Manage Projects Clicked',
    CREATE_NEW_CLICKED = 'Create New Clicked',
    VIEW_DOCS_CLICKED = 'View Docs Clicked',
    VIEW_FORUM_CLICKED = 'View Forum Clicked',
    VIEW_SUPPORT_CLICKED = 'View Support Clicked',
    CREATE_AN_ACCESS_GRANT_CLICKED = 'Create an Access Grant Clicked',
    UPLOAD_USING_CLI_CLICKED = 'Upload Using CLI Clicked',
    UPLOAD_IN_WEB_CLICKED = 'Upload In Web Clicked',
    NEW_PROJECT_CLICKED = 'New Project Clicked',
    LOGOUT_CLICKED = 'Logout Clicked',
    PROFILE_UPDATED = 'Profile Updated',
    PASSWORD_CHANGED = 'Password Changed',
    MFA_ENABLED = 'MFA Enabled',
    BUCKET_CREATED = 'Bucket Created',
    BUCKET_DELETED = 'Bucket Deleted',
    ACCESS_GRANT_CREATED = 'Access Grant Created',
    API_ACCESS_CREATED  = 'API Access Created',
    UPLOAD_FILE_CLICKED = 'Upload File Clicked',
    UPLOAD_FOLDER_CLICKED = 'Upload Folder Clicked',
    CREATE_KEYS_CLICKED = 'Create Keys Clicked',
    DOWNLOAD_TXT_CLICKED = 'Download txt clicked',
    ENCRYPT_MY_ACCESS_CLICKED = 'Encrypt My Access Clicked',
    COPY_TO_CLIPBOARD_CLICKED = 'Copy to Clipboard Clicked',
    CREATE_ACCESS_GRANT_CLICKED = 'Create Access Grant Clicked',
    CREATE_S3_CREDENTIALS_CLICKED = 'Create S3 Credentials Clicked',
    CREATE_KEYS_FOR_CLI_CLICKED = 'Create Keys For CLI Clicked',
    SEE_PAYMENTS_CLICKED = 'See Payments Clicked',
    EDIT_PAYMENT_METHOD_CLICKED = 'Edit Payment Method Clicked',
    USAGE_DETAILED_INFO_CLICKED = 'Usage Detailed Info Clicked',
    ADD_NEW_PAYMENT_METHOD_CLICKED = 'Add New Payment Method Clicked',
    APPLY_NEW_COUPON_CLICKED = 'Apply New Coupon Clicked',
    CREDIT_CARD_REMOVED = 'Credit Card Removed',
    COUPON_CODE_APPLIED = 'Coupon Code Applied',
    INVOICE_DOWNLOADED = 'Invoice Downloaded',
    CREDIT_CARD_ADDED_FROM_BILLING = 'Credit Card Added From Billing',
    STORJ_TOKEN_ADDED_FROM_BILLING = 'Storj Token Added From Billing',
    ADD_FUNDS_CLICKED = 'Add Funds Clicked',
    PROJECT_MEMBERS_INVITE_SENT = 'Project Members Invite Sent',
    UI_ERROR = 'UI error occurred',
    PROJECT_NAME_UPDATED = 'Project Name Updated',
    PROJECT_DESCRIPTION_UPDATED = 'Project Description Updated',
    PROJECT_STORAGE_LIMIT_UPDATED = 'Project Storage Limit Updated',
    PROJECT_BANDWIDTH_LIMIT_UPDATED = 'Project Bandwidth Limit Updated',
}

export enum AnalyticsErrorEventSource {
    ACCESS_GRANTS_PAGE = 'Access grants page',
    ACCOUNT_SETTINGS_AREA = 'Account settings area',
    BILLING_HISTORY_TAB = 'Billing history tab',
    BILLING_COUPONS_TAB = 'Billing coupons tab',
    BILLING_OVERVIEW_TAB = 'Billing overview tab',
    BILLING_PAYMENT_METHODS_TAB = 'Billing payment methods tab',
    BILLING_PAYMENT_METHODS = 'Billing payment methods',
    BILLING_COUPON_AREA = 'Billing coupon area',
    BILLING_APPLY_COUPON_CODE_INPUT = 'Billing apply coupon code input',
    BILLING_PAYMENTS_HISTORY = 'Billing payments history',
    BILLING_PERIODS_SELECTION = 'Billing periods selection',
    BILLING_ESTIMATED_COSTS_AND_CREDITS = 'Billing estimated costs and credits',
    BILLING_ADD_STRIPE_CC_FORM = 'Billing add stripe CC form',
    BILLING_STRIPE_CARD_INPUT = 'Billing stripe card input',
    BILLING_AREA = 'Billing area',
    BILLING_STORJ_TOKEN_CONTAINER = 'Billing STORJ token container',
    BILLING_CC_DIALOG = 'Billing credit card dialog',
    CREATE_AG_MODAL = 'Create access grant modal',
    CONFIRM_DELETE_AG_MODAL = 'Confirm delete access grant modal',
    CREATE_AG_FORM = 'Create access grant form',
    FILE_BROWSER_LIST_CALL = 'File browser - list API call',
    FILE_BROWSER_ENTRY = 'File browser entry',
    PROJECT_INFO_BAR = 'Project info bar',
    CREATE_PROJECT_LEVEL_PASSPHRASE_MODAL = 'Create project level passphrase modal',
    SWITCH_PROJECT_LEVEL_PASSPHRASE_MODAL = 'Switch project level passphrase modal',
    UPGRADE_ACCOUNT_MODAL = 'Upgrade account modal',
    ADD_PROJECT_MEMBER_MODAL = 'Add project member modal',
    ADD_TOKEN_FUNDS_MODAL = 'Add token funds modal',
    CHANGE_PASSWORD_MODAL = 'Change password modal',
    CREATE_PROJECT_MODAL = 'Create project modal',
    DELETE_BUCKET_MODAL = 'Delete bucket modal',
    ENABLE_MFA_MODAL = 'Enable MFA modal',
    DISABLE_MFA_MODAL = 'Disable MFA modal',
    EDIT_PROFILE_MODAL = 'Edit profile modal',
    CREATE_FOLDER_MODAL = 'Create folder modal',
    OBJECT_DETAILS_MODAL = 'Object details modal',
    OPEN_BUCKET_MODAL = 'Open bucket modal',
    SHARE_BUCKET_MODAL = 'Share bucket modal',
    NAVIGATION_ACCOUNT_AREA = 'Navigation account area',
    NAVIGATION_PROJECT_SELECTION = 'Navigation project selection',
    MOBILE_NAVIGATION = 'Mobile navigation',
    BUCKET_CREATION_FLOW = 'Bucket creation flow',
    BUCKET_CREATION_NAME_STEP = 'Bucket creation name step',
    BUCKET_TABLE = 'Bucket table',
    BUCKET_PAGE = 'Bucket page',
    BUCKET_DETAILS_PAGE = 'Bucket details page',
    UPLOAD_FILE_VIEW = 'Upload file view',
    OBJECT_UPLOAD_ERROR = 'Object upload error',
    ONBOARDING_NAME_STEP = 'Onboarding name step',
    ONBOARDING_PERMISSIONS_STEP = 'Onboarding permissions step',
    PROJECT_DASHBOARD_PAGE = 'Project dashboard page',
    PROJECT_USAGE_CONTAINER = 'Project usage container',
    EDIT_PROJECT_DETAILS = 'Edit project details',
    PROJECTS_LIST = 'Projects list',
    PROJECT_MEMBERS_HEADER = 'Project members page header',
    PROJECT_MEMBERS_PAGE = 'Project members page',
    OVERALL_APP_WRAPPER_ERROR = 'Overall app wrapper error',
    OVERALL_GRAPHQL_ERROR = 'Overall graphQL error',
    OVERALL_SESSION_EXPIRED_ERROR = 'Overall session expired error',
    ALL_PROJECT_DASHBOARD = 'All projects dashboard error',
    ONBOARDING_OVERVIEW_STEP = 'Onboarding Overview step error',
    PRICING_PLAN_STEP = 'Onboarding Pricing Plan step error',
}
