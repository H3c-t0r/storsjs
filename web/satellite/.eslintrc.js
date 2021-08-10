// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

module.exports = {
    root: true,
    env: {
        node: true
    },
    extends: [
        'plugin:vue/essential',
        'eslint:recommended',
        '@vue/typescript/recommended',
    ],
    parserOptions: {
        ecmaVersion: 2020
    },
    rules: {
        'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
        'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',

        "indent": ["warn", 4],

        "@typescript-eslint/no-unused-vars": [
            "warn", {
                "vars": "all",
                "args": "all",
                "argsIgnorePattern": "^_"
            }],

        '@typescript-eslint/no-empty-function': "off",
        '@typescript-eslint/no-var-requires': "off",

        '@typescript-eslint/no-explicit-any': "off" // TODO: not everything has been fixed yet
    },
}