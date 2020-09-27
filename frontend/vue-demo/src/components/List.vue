<template>
  <table class="table">
    <thead>
      <tr>
        <th>ID</th>
        <th>タスク名</th>
        <th>メモ</th>
        <th>期日</th>
        <th>完了</th>
        <th>削除</th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="todo in todos" v-bind:key="todo.id">
        <td>{{ todo.id }}</td>
        <td>{{ todo.title }}</td>
        <td>{{ todo.note }}</td>
        <td>{{todo.due_date }}</td>
        <td>
          <div class="text-center">
            <button class="btn" v-on:click="callCompleteTodo(todo)" v-bind:class="{'btn-primary': !todo.is_completed, 'btn-secondary': todo.is_completed}">
              <span v-if="todo.is_completed">
                済み
              </span>
              <span v-if="!todo.is_completed">
                完了する
              </span>
            </button>
          </div>
        </td>
        <td>
          <div class="text-center">
            <button class="btn btn-danger" v-on:click="callDeleteTodo(todo.id)">削除</button>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script>
export default {
  props: ['todos', 'apiError',],
  methods: {
    callDeleteTodo(id) {
      this.$emit('deleteTodo', id)
    },
    callCompleteTodo(todo) {
      this.$emit('completeTodo', todo)
    },
  },
  created() {
    this.$emit('getTodo')
  }
}
</script>