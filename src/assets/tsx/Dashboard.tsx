import { FC } from 'react';

type DashboardIconProps = {
  className?: string;
  fill: string;
};
export const DashboardIcon: FC<DashboardIconProps> = ({ fill, className }) => (
  <>
    <svg fill={fill} className={className} width='800px' height='800px' viewBox='0 0 24 24' id='dashboard'>
      <rect x='2' y='2' width='9' height='11' rx='2'></rect>
      <rect x='13' y='2' width='9' height='7' rx='2'></rect>
      <rect x='2' y='15' width='9' height='7' rx='2'></rect>
      <rect x='13' y='11' width='9' height='11' rx='2'></rect>
    </svg>
  </>
);
