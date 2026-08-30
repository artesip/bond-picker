import { AgGridReact } from 'ag-grid-react';
import { AllCommunityModule  } from 'ag-grid-community';

import { shadcnTheme } from '#/ag/theme';
import { usePickedBonds } from '#/entities/bonds/hooks';
import { withRuLocalization } from '#/ag/localization';
import { DeleteIcon, RatingIcon } from '#/components/cell-renders/components';

import { coldef } from './col-def';
import { RatingDrawer } from './ui/rating-drawer';

import type { GetRowIdParams } from 'ag-grid-community';

export function ChosenPage() {
  const { data } = usePickedBonds();

  return (
    <div className='h-full'>
      <AgGridReact
        rowData={ data }
        columnDefs={ coldef }
        modules={ [AllCommunityModule] }
        theme={ shadcnTheme }
        gridOptions={ { ...withRuLocalization() } }
        enableCellTextSelection={ true }
        rowSelection={ { mode: 'multiRow' } }
        components={ {
          deleteIcon: DeleteIcon,
          ratingIcon: RatingIcon,
        } }
        getRowId={ (params: GetRowIdParams) => String(params.data.id) }
      /> 

      <RatingDrawer bonds={ data || [] }/>
    </div>
  );
}