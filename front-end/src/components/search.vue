<template>
  <div>
    <sun-header></sun-header>
  </div>
</template>

<script>
import SunHeader from './base-components/SunHeader'
export default {
  components: {
    SunHeader
  },
  data () {
    return {
      artworks: []

    }
  },
  name: 'Search',
  methods: {

  }
}
</script>
