import firebase from 'firebase/app';
import firebaseConfig from 'config/firebase-config';

const init = () => firebase.initializeApp(firebaseConfig);

export default {
  init,
};
