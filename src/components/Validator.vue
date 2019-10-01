<template>
  <validation-provider
    v-slot="{ errors }"
    :rules="rulesText"
    :name="name"
    :mode="mode">
    <slot />
    <span class="text-danger text-sm">
      {{ errors[0] }}
    </span>
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
      default: 'passive',
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
