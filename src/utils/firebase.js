import firebase from 'firebase/app';
import firebaseConfig from 'config/firebase-credential';

const init = () => firebase.initializeApp(firebaseConfig);

export default {
  init,
};
