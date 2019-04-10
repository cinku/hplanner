<template>
  <div class="planner">
    <logo />
    <div v-if="!searching">
      <Destination v-if="!showTrip" @submitLocation="handleSubmit"/>
      <trip-details :trip-data="tripData" v-else/>
    </div>
    <Preloader v-else/>
    <div v-if="searching || showTrip">
      <small><a href="#" class="cancel-link" v-on:click="cancel">Anuluj</a></small>
    </div>
  </div>
</template>

<script>
import Destination from './Destination.vue';
import Preloader from '../shared/Preloader.vue';
import TripDetails from './TripDetails.vue';
import Logo from '../shared/Logo.vue';

export default {
  components: {
    Destination,
    Preloader,
    TripDetails,
    Logo,
  },
  props: ['showTrip', 'tripData', 'searching'],
  methods: {
    handleSubmit(formData) {
      this.$emit('submitLocation', formData);
    },
    cancel() {
      this.$emit('cancel');
    },
  },
};
</script>

<style lang="scss" scoped>
  .planner {
    padding: 40px;
    background: #FFFFFF;
    min-height: 100vh;
  }
  .cancel-link {
    text-decoration: underline;
    color: #55b0f3;
  }
</style>
