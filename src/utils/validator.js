import { localize, extend } from 'vee-validate';
import { required, min_value as minValue, max_value as maxValue } from 'vee-validate/dist/rules';
import ko from 'vee-validate/dist/locale/ko.json';
import isURL from 'validator/lib/isURL';

// [참고] 기본 메시지 양식
// https://github.com/logaretm/vee-validate/blob/master/locale/ko.json

localize('ko', ko);
extend('required', {
  ...required,
  message: '{_field_} 항목은 필수 정보입니다',
});

extend('minValue', {
  ...minValue,
  message: '{_field_}항목의 값은 {min} 이상이어야 합니다',
});

extend('maxValue', {
  ...maxValue,
  message: '{_field_}항목의 값은 {max} 이하이어야 합니다',
});

extend('url', {
  validate: (value) => isURL(value),
  message: '{_field_} 항목의 값은 유효한 URL 형식이어야 합니다',
});

export { ValidationProvider, ValidationObserver } from 'vee-validate';
