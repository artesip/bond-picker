import js from '@eslint/js';
import globals from 'globals';
import reactHooks from 'eslint-plugin-react-hooks';
import reactRefresh from 'eslint-plugin-react-refresh';
import tseslint from 'typescript-eslint';
import jsxA11y from 'eslint-plugin-jsx-a11y';
import pluginReact from 'eslint-plugin-react';
import tsParser from '@typescript-eslint/parser';
import reactCompiler from 'eslint-plugin-react-compiler';
import eslintPluginImportX from 'eslint-plugin-import-x';
import { createTypeScriptImportResolver } from 'eslint-import-resolver-typescript';

export default tseslint.config(
  {
    // Глобальное игнорирование
    ignores: ['dist', '.next/**'],
  },
  js.configs.recommended,
  ...tseslint.configs.recommended,
  pluginReact.configs.flat.recommended,
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
      'jsx-a11y'      : jsxA11y,
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
      react: {
        version: 'detect',
      },
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

      // Правила форматирования (Внимание: в ESLint 10 они могут потребовать @stylistic)
      'key-spacing'                        : ['error', { align: 'colon' }],
      'jsx-quotes'                         : ['error', 'prefer-single'],
      'keyword-spacing'                    : ['error', { before: true, after: true }],
      'react/jsx-indent'                   : ['error', 2],
      'react/jsx-indent-props'             : ['error', 2],
      'object-curly-spacing'               : ['error', 'always'],
      'semi'                               : ['error', 'always', { omitLastInOneLineBlock: true }],
      'react/destructuring-assignment'     : 'off',
      'react/no-unstable-nested-components': ['error', { allowAsProps: true }],
      'linebreak-style'                    : ['error', 'unix'],
      'react/react-in-jsx-scope'           : 'off',
      'react/jsx-props-no-spreading'       : 'off',
      'max-depth'                          : ['error', { max: 6 }],
      'react/jsx-curly-spacing'            : [2, 'always'],
      'react/jsx-filename-extension'       : [1, { extensions: ['.js', '.jsx', '.tsx'] }],
      'indent'                             : ['error', 2, { ignoredNodes: ['JSXAttribute'] }],
      'react/require-default-props'        : [0],

      // Accessibility
      'jsx-a11y/anchor-is-valid'               : 1,
      'jsx-a11y/label-has-associated-control'  : 1,
      'jsx-a11y/no-static-element-interactions': 1,

      'comma-dangle' : ['error', 'only-multiline'],
      'comma-spacing': ['error', { before: false, after: true }],
      
      // ИСПРАВЛЕНО: оператор переноса строки
      'operator-linebreak': ['error', 'before', { overrides: { '?': 'before', ':': 'before' } }],
      
      'quotes'             : ['error', 'single'],
      'space-before-blocks': 'error',
      'space-infix-ops'    : 'error',
      
      // ИСПРАВЛЕНО: заместо удаленного no-new-object
      'no-object-constructor': 'error',
      'no-lonely-if'         : 'error',
    },
  }
);