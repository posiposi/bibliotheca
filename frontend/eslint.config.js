import { Linter } from "eslint";
import eslintRecommended from "eslint/conf/eslint-recommended";
import reactRecommended from "eslint-plugin-react";
import typescriptRecommended from "@typescript-eslint/eslint-plugin/dist/configs/recommended";

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
  override: [eslintRecommended, reactRecommended, typescriptRecommended],
};

export default config;
