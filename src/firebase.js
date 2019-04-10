import firebase from 'firebase/app';
import 'firebase/firestore';

firebase.initializeApp({
  projectId: '',
  databaseURL: '',
});

const db = firebase.firestore();

function getRequestsCollection() {
  return db.collection('requests');
}

function getRequestById(referenceId) {
  return getRequestsCollection().doc(referenceId);
}

function getRequestsForCity(city) {
  return getRequestsCollection().where('city', '==', city);
}

export {
  getRequestById,
  getRequestsCollection,
  getRequestsForCity,
};
