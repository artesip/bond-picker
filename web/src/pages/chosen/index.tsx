import { AgGridReact } from 'ag-grid-react';
import { AllCommunityModule  } from 'ag-grid-community';


import { shadcnTheme } from '#/ag/theme';
import { usePickedBonds } from '#/entities/bonds/hooks';
import { withRuLocalization } from '#/ag/localization';
import { DeleteIcon, RatingIcon } from '#/components/cell-renders/components';
import { DrawerTrigger,  DrawerContent,  DrawerHeader,  DrawerTitle,  DrawerDescription,  DrawerFooter,  DrawerClose, Drawer } from '#/components/ui/drawer';
import { Button } from '#/components/ui/button';

import { coldef } from './col-def';

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


      <Drawer direction='right'>
        <DrawerTrigger>Open</DrawerTrigger>
        <DrawerContent>
          <DrawerHeader>
            <DrawerTitle>Are you absolutely sure?</DrawerTitle>
            <DrawerDescription>This action cannot be undone.</DrawerDescription>
          </DrawerHeader>
          <DrawerFooter>
            <Button>Submit</Button>
            <DrawerClose>
              <Button variant='outline'>Cancel</Button>
            </DrawerClose>
          </DrawerFooter>
        </DrawerContent>
      </Drawer>
    </div>
  );
}