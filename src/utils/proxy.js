import { chain, isNull, has, isEqual, includes } from 'lodash';

export const makeProxy = (obj) => {
  const targetObj = {
    original: obj,
    diff(includeBefore = false) {
      return chain(this)
        .mapValues((val, key) => {
          if ((!has(this.original, key)) || (has(this, key) && !isEqual(val, this.original[key]))) {
            if (!includeBefore) {
              return this[key];
            } else {
              return {
                before: val,
                after: this[key],
              };
            }
          }
          return null;
        })
        .pickBy((val) => !isNull(val))
        .value();
    },
  };
  return new Proxy(targetObj, {
    get: (target, key, receiver) => {
      if (has(target, key) && !isNull(target[key])) {
        return Reflect.get(target, key, receiver);
      }
      return Reflect.get(obj, key);
    },
    set: (target, key, value, receiver) => Reflect.set(target, key, value, receiver),
    ownKeys: (target) => {
      return chain(target)
        .keys()
        .filter((val) => !includes(['original', 'diff'], val))
        .value();
    },
  });
};
