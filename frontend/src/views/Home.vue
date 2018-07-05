<template>
  <div class="home">
    <businesscard msg="Визитки"/>
  </div>
</template>

<script>
// @ is an alias to /src
import businesscard from '@/components/businesscard.vue'

export default {
  name: 'home',
  components: {
    businesscard
  }
}
</script>
