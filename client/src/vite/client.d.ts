interface ImportMetaEnv {
  VITE_WORD_LENGTH: string;
  VITE_MAX_ATTEMPTS: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
