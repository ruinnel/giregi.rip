<template>
  <validation-provider
    v-slot="{ errors }"
    :rules="rulesText"
    :name="name"
    :mode="mode"
    :slim="slim"
  >
    <slot />
    <div class="invalid-feedback">{{ errors[0] }}</div>
  </validation-provider>
</template>

<script>
import { join, isArray, includes } from 'lodash';
import { ValidationProvider } from 'vee-validate';
export default {
  name: 'Validator',
  component: {
    ValidationProvider,
  },
  props: {
    name: {
      type: String,
      required: true,
    },
    rules: {
      type: String,
      default: 'required',
    },
    mode: {
      type: String,
      validator: (val) => includes(['aggressive', 'passive', 'lazy', 'eager'], val),
      default: 'aggressive',
    },
    slim: {
      type: Boolean,
      default: true,
    },
  },
  computed: {
    rulesText() {
      if (isArray(this.rules)) {
        return join(this.rules, '|');
      }
      return this.rules;
    },
  },
};
</script>
