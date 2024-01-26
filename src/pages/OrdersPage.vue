<script setup lang="ts">
import {Search} from '@element-plus/icons-vue';
import {refDebounced} from '@vueuse/core';
import {computed, Ref, ref, watchEffect} from 'vue';

import {Order} from '../types/order';
import {Response} from '../types/response';
import {Server} from '../utils/axios';

const keyword = ref('');
const keywordDebounced = refDebounced(keyword, 1000);

const filterDate = ref('');
const loading = ref(true);
const currentPage = ref(1);
const pageSize = ref(5);
const totalData = ref(0);
const orders: Ref<Order[]> = ref([] satisfies Order[]);

const queryParams = computed(() => {
  return new URLSearchParams({
    page: String(currentPage.value),
    limit: String(pageSize.value),
    keyword: keywordDebounced.value,
    start_date: '',
    end_date: '',
  }).toString();
});

const fetchOrders = async (): Promise<void> => {
  try {
    const res = await Server.get<Response<Order[]>>(
      `api/orders?${queryParams.value}`,
    );

    orders.value = res.data.result;
    totalData.value = res.data.meta?.total ?? 0;
  } catch (error) {
    console.error('error fetching orders:', error);
  } finally {
    loading.value = false;
  }
};

watchEffect(() => void fetchOrders());

const handleSizeChange = (val: number): void => {
  pageSize.value = val;
};
const handleCurrentChange = (val: number): void => {
  currentPage.value = val;
};
</script>

<template>
  <div>
    <!-- Search Bar -->
    <div class="mb-3">
      <el-input
        v-model="keyword"
        placeholder="Input your keyword"
        size="large"
        class="input-with-select"
      >
        <template #prepend>
          <el-button
            type="info"
            :icon="Search"
            >Search</el-button
          >
        </template>
      </el-input>
    </div>

    <!-- Filter Date -->
    <div class="filter-date-range">
      <span class="text-secondary">Created Date</span>
      <el-date-picker
        v-model="filterDate"
        type="daterange"
        range-separator="-"
        start-placeholder="Start date"
        end-placeholder="End date"
        size="large"
      />
    </div>

    <!-- Total Amount -->
    <span class="text-secondary">Total amount: <strong>$198.23</strong></span>

    <!-- Table -->
    <div class="mt-5">
      <el-table
        v-loading="loading"
        :data="orders"
        style="width: 100%"
      >
        <el-table-column
          label="Order Name"
          prop="order_name"
        />
        <el-table-column
          label="Customer Company"
          prop="customer_company_name"
        />
        <el-table-column
          label="Customer Name"
          prop="customer_name"
        />
        <el-table-column
          label="Order Date"
          prop="order_date"
        />
        <el-table-column
          label="Delivered Amount"
          prop="delivered_amount"
        />
        <el-table-column
          label="Total Amount"
          prop="total_amount"
        />
      </el-table>
    </div>

    <!-- Pagination -->
    <div class="pagination">
      <div class="pagination-block">
        <el-pagination
          :current-page="currentPage"
          :page-size="pageSize"
          :page-sizes="[5, 10, 20, 30]"
          :background="true"
          layout="total, sizes, prev, pager, next, jumper"
          :total="totalData"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.input-with-select .el-input-group__prepend {
  background-color: var(--el-fill-color-blank);
}

.filter-date-range {
  padding: 20px 0;
  flex: 1;
}

.filter-date-range span {
  display: block;
  margin-bottom: 5px;
}

.text-secondary {
  color: var(--el-text-color-secondary);
  font-size: 16px;
}

.pagination {
  display: flex;
  justify-content: flex-end;
  width: 100%;
  padding: 0;
  flex-wrap: wrap;
}

.pagination-block {
  margin: 30px 0;
  text-align: center;
}
</style>
