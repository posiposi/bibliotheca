import { Linter } from "eslint";

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
  plugins: ["react"],
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
