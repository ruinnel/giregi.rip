import { map } from 'lodash';

const labels = {
  url: { name: 'URL', icon: 'fas fa-globe-asia' },
  title: { name: '제목', icon: 'far fa-newspaper' },
  writer: { name: '작성자', icon: 'fas fa-user-edit' },
  createdAt: { name: '작성일', icon: 'far fa-clock' },
  updatedAt: { name: '수정일', icon: 'fas fa-clock' },
  email: { name: '이메일', icon: 'far fa-envelope' },
  agency: { name: '언론사', icon: 'far fa-building' },
  cowriter: { name: '공동 작성자', icon: 'fas fa-user-friends' },
};

const convert = (archive, vm) => {
  return map(archive.summary, ({ name, value }) => {
    let converted = value;
    if (name === 'createdAt' || name === 'updatedAt') {
      converted = vm.formatDateTime(value);
    }
    return { ...labels[name], value: converted };
  });
};

export default {
  labels,
  convert,
};
