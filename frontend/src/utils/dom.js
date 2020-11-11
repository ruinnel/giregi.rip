export default {
  insertBody(element, parent) {
    const body = parent || document.body;
    body.insertBefore(element, body.firstChild);
  },
  removeBody(element, parent) {
    const body = parent || document.body;
    body.removeChild(element);
  },
};
