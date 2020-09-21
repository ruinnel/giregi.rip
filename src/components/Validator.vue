<template>
  <validation-provider
    v-slot="{ errors }"
    :rules="rulesText"
    :name="name"
    :mode="mode"
    :slim="slim"
  >
    <slot />
    <div class="invalid-feedback" :style="getStyle(errors[0])">{{ errors[0] }}</div>
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
  methods: {
    getStyle(e) {
      if (!this.$el) return {};
      const formControl = this.$el.querySelector('.form-control');
      if (e) {
        formControl.classList.add('is-invalid');
        return { display: 'block' };
      } else {
        formControl.classList.remove('is-invalid');
        return { display: 'none' };
      }
    },
  },
};
</script>
