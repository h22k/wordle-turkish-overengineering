interface ImportMetaEnv {
    VITE_WORD_LENGTH: string;
    VITE_MAX_ATTEMPTS: string;
    VITE_API_BASE_URL: string;
}

interface ImportMeta {
    readonly env: ImportMetaEnv;
}
