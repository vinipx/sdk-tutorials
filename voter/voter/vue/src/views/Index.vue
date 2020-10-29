<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form type="poll" :fields="['title', 'options', ]" module="voter" />
      <poll-form />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";
import PollForm from "@/components/PollForm"

export default {
  components: { ...sp, PollForm },
};
</script>
