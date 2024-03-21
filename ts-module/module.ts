import _ from "lodash";

export function greet(name: string): string {
  return `hello, ${_.upperFirst(name)}!`;
}
