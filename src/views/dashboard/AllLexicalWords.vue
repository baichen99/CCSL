<template>
  <div class="app-container">WORDS</div>
</template>

<script>
import {
  GetLexicalWordsList,
  CreateLexicalWords,
  UpdateLexicalWords,
  DeleteLexicalWords
} from "@/api/words";

export default {
  name: "AllLexicalWords",
  data() {
    return {
      params: {
        initial: "",
        chinese: "",
        english: "",
        pos: ""
      }
    };
  }
};
</script>