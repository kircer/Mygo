<template>
  <main class="page">
    <h1>User List From MySQL</h1>

    <button class="btn" @click="fetchUsers" :disabled="loading">
      {{ loading ? 'Loading...' : 'Refresh Users' }}
    </button>

    <p v-if="error" class="error">{{ error }}</p>

    <table v-if="users.length" class="table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Age</th>
          <th>Address</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.id }}</td>
          <td>{{ user.name }}</td>
          <td>{{ user.age }}</td>
          <td>{{ user.address }}</td>
        </tr>
      </tbody>
    </table>

    <p v-else-if="!loading" class="empty">No users found.</p>
  </main>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

type User = {
  id: string
  name: string
  age: number
  address: string
}

type UsersResponse = {
  data: User[]
  error?: string
}

const users = ref<User[]>([])
const loading = ref(false)
const error = ref('')

async function fetchUsers() {
  loading.value = true
  error.value = ''

  try {
    const response = await fetch('/api/users')
    const payload = (await response.json()) as UsersResponse

    if (!response.ok) {
      throw new Error(payload.error || 'Request failed')
    }

    users.value = payload.data || []
  } catch (e) {
    const message = e instanceof Error ? e.message : String(e)
    error.value = `Failed to load users: ${message}. Make sure test.go is running on :8080.`
  } finally {
    loading.value = false
  }
}

onMounted(fetchUsers)
</script>

<style scoped>
.page {
  max-width: 860px;
  margin: 40px auto;
  padding: 16px;
  font-family: "Segoe UI", "Microsoft YaHei", sans-serif;
}

h1 {
  margin-bottom: 16px;
}

.btn {
  margin-bottom: 16px;
  padding: 8px 14px;
  border: none;
  border-radius: 6px;
  background: #0f766e;
  color: #fff;
  cursor: pointer;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.error {
  color: #b91c1c;
  margin-bottom: 12px;
}

.empty {
  color: #475569;
}

.table {
  width: 100%;
  border-collapse: collapse;
}

.table th,
.table td {
  border: 1px solid #d1d5db;
  text-align: left;
  padding: 8px;
}

.table th {
  background: #f1f5f9;
}
</style>