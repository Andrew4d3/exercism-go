export default class HelloWorld {
  static hello(name: string = "World"): string {
    return `Hello, ${name}!`;
  }
}

export const hello = HelloWorld.hello;
