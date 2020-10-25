import { isNumber } from 'lodash';
import { DateTime } from 'luxon';

const numberFormat = Intl.NumberFormat('ko-KR', { style: 'decimal', maximumFractionDigits: 0 });

const formatNumber = (num) => {
  if (!num) { // 0 || undefined || null || NaN
    return numberFormat.format(0);
  }
  return numberFormat.format(num);
};

const formatDate = (date, format = 'yyyy-MM-dd') => {
  if (isNumber(date)) {
    return DateTime.fromMillis(date).toFormat(format);
  } else if (date instanceof DateTime) {
    return date.toFormat(format);
  } else if (date instanceof Date) {
    return DateTime.fromJSDate(date).toFormat(format);
  } else {
    return '';
  }
};

const formatDateTime = (date, format = 'yyyy-MM-dd HH:mm:ss') => {
  return formatDate(date, format);
};

const formatTime = (inputSeconds) => {
  const value = parseInt(inputSeconds, 10);

  let hours = Math.floor(value / 3600);
  let minutes = Math.floor((value - (hours * 3600)) / 60);
  let seconds = value - (hours * 3600) - (minutes * 60);

  if (hours < 10) {
    hours = `0${hours}`;
  }
  if (minutes < 10) {
    minutes = `0${minutes}`;
  }
  if (seconds < 10) {
    seconds = `0${seconds}`;
  }

  return `${hours}:${minutes}:${seconds}`;
};

export default {
  formatNumber,
  formatDate,
  formatDateTime,
  formatTime,
};
