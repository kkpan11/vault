/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: BUSL-1.1
 */

import config from 'vault/config/environment';

const { DEFAULT_PAGE_SIZE } = config.APP;

type PaginatedData = string | Record<string, string>[];
type PaginateOptions = {
  page?: number;
  pageSize?: number;
  filter?: string;
  filterKey?: string;
};

/**
 * Util to paginate data set based on page number and size
 * If filter is provided, it will filter the data prior to paginating
 */

export const paginate = (data: PaginatedData, options: PaginateOptions = {}) => {
  const { page = 1, pageSize = DEFAULT_PAGE_SIZE, filter, filterKey } = options;
  let filteredData = data;

  if (Array.isArray(data)) {
    // filter data before paginating if filter is provided
    if (filter) {
      filteredData = data.filter((item) => {
        const filterValue = filterKey ? item[filterKey] : item;
        if (typeof filterValue === 'string') {
          return filterValue.toLowerCase().includes(filter.toLowerCase());
        }
        return false;
      });
    }

    const lastPage = Math.ceil(filteredData.length / pageSize);
    const start = (page - 1) * pageSize;
    const end = start + pageSize;
    filteredData = filteredData.slice(start, end);
    // add meta data previously from lazyPaginatedQuery since components expect it
    Object.defineProperty(filteredData, 'meta', {
      value: {
        currentPage: page,
        lastPage,
        nextPage: page + 1,
        prevPage: page - 1,
        total: data.length,
        filteredTotal: filteredData.length,
        pageSize,
      },
    });
  }

  return filteredData;
};
