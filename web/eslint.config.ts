import js from '@eslint/js';
import globals from 'globals';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import tseslint from 'typescript-eslint';
import tsParser from '@typescript-eslint/parser';
import reactCompiler from 'eslint-plugin-react-compiler';
import eslintPluginImportX from 'eslint-plugin-import-x';
import { createTypeScriptImportResolver } from 'eslint-import-resolver-typescript';

export default tseslint.config(
  {
    // Глобальное игнорирование
    ignores: [
      'node_modules/**',
      'dist/**',
      '.output/**',
      '.next/**',
      '.tanstack/**',
      '.nitro/**',
      '.nuxt/**',
      'coverage/**',
      'public/**',
      'build/**',
      'vite.config.ts',
    ],
  },
  js.configs.recommended,
  ...tseslint.configs.recommended,
  eslintPluginImportX.flatConfigs.recommended,
  eslintPluginImportX.flatConfigs.typescript,
  {
    files          : ['**/*.{ts,tsx}'],
    languageOptions: {
      parser     : tsParser,
      sourceType : 'module',
      ecmaVersion: 'latest',
      globals    : {
        ...globals.browser,
      },
    },
    plugins: {
      'react-hooks'   : reactHooks,
      'react-refresh' : reactRefresh,
      'react-compiler': reactCompiler,
    },
    settings: {
      'import-x/resolver-next': [
        createTypeScriptImportResolver({
          alwaysTryTypes: true,
          project       : './tsconfig.app.json',
        }),
      ],
    },
    rules: {
      ...reactHooks.configs.recommended.rules,
      'react-refresh/only-export-components': [
        'warn',
        { allowConstantExport: true },
      ],

      'react-compiler/react-compiler'    : 'error',
      'no-unused-vars'                   : 'off',
      '@typescript-eslint/no-unused-vars': 'warn', // Рекомендую включить TS версию вместо выключенной JS

      'import-x/no-dynamic-require'             : 'warn',
      'import-x/no-nodejs-modules'              : 'warn',
      'import-x/consistent-type-specifier-style': ['error', 'prefer-top-level'],
      'import-x/order'                          : [
        'error',
        {
          'newlines-between': 'always',
          groups            : ['builtin', 'external', 'internal', 'parent', 'sibling', 'index', 'object', 'type'],
        },
      ],

      '@typescript-eslint/consistent-type-imports': ['error', { prefer: 'type-imports' }],
      '@typescript-eslint/no-empty-function'      : 'off',

      // Правила форматирования
      'key-spacing'         : ['error', { align: 'colon' }],
      'keyword-spacing'     : ['error', { before: true, after: true }],
      'object-curly-spacing': ['error', 'always'],
      'semi'                : ['error', 'always', { omitLastInOneLineBlock: true }],
      'max-depth'           : ['error', { max: 6 }],

      'jsx-quotes': ['error', 'prefer-single'],

      'comma-dangle' : ['error', 'only-multiline'],
      'comma-spacing': ['error', { before: false, after: true }],

      'operator-linebreak': ['error', 'before', { overrides: { '?': 'before', ':': 'before' } }],

      'quotes'             : ['error', 'single'],
      'space-before-blocks': 'error',
      'space-infix-ops'    : 'error',

      'no-object-constructor': 'error',
      'no-lonely-if'         : 'error',
    },
  }
);
