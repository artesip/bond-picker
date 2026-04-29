import type { ColDef } from 'ag-grid-community';

export const coldef: ColDef[] = [
  {
    field     : 'name',
    headerName: 'Имя',
    flex      : 1,
    filter    : 'agTextColumnFilter',
    minWidth  : 150,
  },
  {
    field     : 'price',
    headerName: 'Цена',
    filter    : 'agNumberColumnFilter',
    flex      : 1,
    minWidth  : 100,
  },
  {
    field     : 'ytm',
    headerName: 'Доходность',
    filter    : 'agNumberColumnFilter',
    flex      : 1,
    minWidth  : 100,
  },
  {
    field     : 'duration',
    headerName: 'Дюрация',
    filter    : 'agNumberColumnFilter',
    flex      : 1,
    minWidth  : 100,
  },
  {
    field     : 'couponPercent',
    headerName: 'Купон',
    filter    : 'agNumberColumnFilter',
    flex      : 1,
    minWidth  : 100,
  },
  {
    field         : 'matDate',
    headerName    : 'Погашение',
    filter        : 'agDateColumnFilter',
    flex          : 1,
    minWidth      : 100,
    valueFormatter: (params) => {
      if (!params.value) return '';

      const date = new Date(params.value);

      const day = String(date.getDate()).padStart(2, '0');
      const month = String(date.getMonth() + 1).padStart(2, '0');
      const year = date.getFullYear();

      return `${day}.${month}.${year}`;
    },
  },
  {
    field     : 'acruedint',
    headerName: 'НКД',
    filter    : 'agNumberColumnFilter',
    flex      : 1 / 2,
    minWidth  : 100,
  },
  {
    field     : 'count',
    headerName: 'Кол-во',
    filter    : 'agNumberColumnFilter',
    flex      : 1 / 2,
    minWidth  : 80,
  },
  {
    headerName  : 'Рейтинг',
    cellRenderer: 'ratingIcon',
    flex        : 1 / 2,
    sortable    : false,
    minWidth    : 80,
    cellClass   : 'center-header'
  },
  {
    headerName     : 'Действия',
    headerComponent: 'deleteIcon',
    flex           : 1 / 4,
    sortable       : false,
    minWidth       : 60,
    headerClass    : 'center-header',
  },
];