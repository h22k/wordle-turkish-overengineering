// @ts-check

import eslint from '@eslint/js'
import tseslint from 'typescript-eslint'
import prettier from 'eslint-plugin-prettier'
import eslintConfigPrettier from 'eslint-config-prettier'

export default tseslint.config( 
  {
    files: [ '**/*.{ts,tsx}' ],
    plugins: {
      prettier: prettier,
    },
    rules: {
      'prettier/prettier': [ 'error', {
        semi: false,
        singleQuote: true,
        trailingComma: 'es5',
        bracketSpacing: true
      }],
      'react/prop-types': 'off',
      'react/react-in-jsx-scope': 'off',
      'no-unused-vars': 'warn',
      '@typescript-eslint/no-unused-vars': 'warn',
      '@typescript-eslint/explicit-function-return-type': 'off',
      '@typescript-eslint/explicit-module-boundary-types': 'off',
      '@typescript-eslint/no-explicit-any': 'warn',
      'max-len': [ 'error', { code: 100 } ],
      'no-console': 'warn',
      'no-debugger': 'warn',
      'space-in-parens': [ 'error', 'always' ],
      'array-bracket-spacing': [ 'error', 'always' ],
      'object-curly-spacing': [ 'error', 'always' ],
      'template-curly-spacing': [ 'error', 'always' ]
    },
    settings: {
      react: {
        version: 'detect',
      },
    },
  },
  eslint.configs.recommended,
  ...tseslint.configs.recommended,
  eslintConfigPrettier
)
