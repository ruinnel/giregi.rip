export default (methods, client) => new Proxy(
  methods,
  {
    get(target, propKey, receiver) {
      const origMethod = target[propKey];
      return function (...args) {
        return origMethod.apply(this, [client, ...args]);
      };
    },
  },
);
