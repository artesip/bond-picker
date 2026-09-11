import { Star, MousePointerClick, Percent } from 'lucide-react';

export const buttons = {
  base: [
    {
      icon : <Star />,
      url  : '/app/chosen',
      title: 'Избранное',
    },
    {
      icon : <MousePointerClick/>,
      url  : '/app/picker',
      title: 'Выбор облигаций',
    },
  ],
  analitics: [
    {
      icon : <Percent/>,
      url  : '/app/key-rate',
      title: 'Ключевая ставка',
    },
  ]
};