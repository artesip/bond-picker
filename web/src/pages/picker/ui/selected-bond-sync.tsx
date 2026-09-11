import { useEffect } from 'react';
import { useSearch } from '@tanstack/react-router';

import { useSelectedBondStore } from '#/stores/selected-bond';

type SelectedBondSyncProps = {
  from: '/app/picker' | '/app/watch'
};

export function SelectedBondSync({ from }: SelectedBondSyncProps) {
  const { id } = useSearch({ from });
  const setSelectedId = useSelectedBondStore((state) => state.setSelectedId);

  useEffect(() => {
    const current = useSelectedBondStore.getState().selectedId;

    if (id === undefined && current !== null) {
      setSelectedId(null);
    } else if (current === null && id !== undefined) {
      setSelectedId(id);
    }
  }, [id, setSelectedId]);

  useEffect(() => {
    const handlePopState = () => {
      setSelectedId(new URLSearchParams(window.location.search).get('id'));
    };

    window.addEventListener('popstate', handlePopState);

    return () => window.removeEventListener('popstate', handlePopState);
  }, [setSelectedId]);

  return null;
}
