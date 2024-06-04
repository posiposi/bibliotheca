import { Linter } from "eslint";
import pluginObject from "react";

/** @type {Linter.Config} */
const config = {
  files: ["**/*.js", "**/*.ts", "**/*.jsx", "**/*.tsx"],
  languageOptions: {
    ecmaVersion: 12,
    sourceType: "module",
    globals: {
      Atomics: "readonly",
      SharedArrayBuffer: "readonly",
    },
  },
  plugins: {
    react: pluginObject,
  },
  settings: {
    react: {
      version: "detect",
    },
  },
  rules: {
    "react/react-in-jsx-scope": "off",
  },
  extends: [
    "eslint:recommended",
    "plugin:react/recommended",
    "plugin:@typescript-eslint/recommended",
  ],
};

export default config;
