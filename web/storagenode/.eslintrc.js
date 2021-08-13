// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

module.exports = {
    root: true,
    env: {
        node: true
    },
    extends: [
        'plugin:vue/recommended',
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
        "vue/html-indent": ["warn", 4],

        "@typescript-eslint/no-explicit-any": "off", // TODO: fix

        "@typescript-eslint/no-unused-vars": [
            "warn", {
                "vars": "all",
                "args": "all",
                "argsIgnorePattern": "^_"
            }],

        '@typescript-eslint/no-empty-function': "off",
        '@typescript-eslint/no-var-requires': "off",

        "vue/max-attributes-per-line": ["off"],
        "vue/singleline-html-element-content-newline": ["off"],

        "vue/no-v-html": ["off"], // needs a dedicated fix

        "vue/block-lang": ["error", {"script": {"lang": "ts"}}],
        "vue/html-button-has-type": ["error"],
        "vue/no-unused-properties": ["warn"],
        "vue/no-unused-refs": ["warn"],
        "vue/no-useless-v-bind": ["warn"],
    },
}