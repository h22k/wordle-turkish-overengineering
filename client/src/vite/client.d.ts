interface ImportMetaEnv {
  VITE_WORD_LENGTH: string;
  VITE_MAX_ATTEMPTS: string;
  VITE_API_URL_LOCAL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
