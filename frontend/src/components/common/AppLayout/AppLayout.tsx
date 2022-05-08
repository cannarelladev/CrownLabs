import { InfoOutlined } from '@ant-design/icons';
import { Col, Layout, Progress, Result, Row, Tooltip } from 'antd';
import { FC, useState } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { logout } from '../../../contexts/AuthContext';
import { PUBLIC_URL } from '../../../env';
import { RouteDescriptor } from '../../../utils';
import Navbar from '../Navbar';
import SidebarInfo from '../SidebarInfo';
import TooltipButton from '../TooltipButton';
import { TooltipButtonData } from '../TooltipButton/TooltipButton';
import './AppLayout.less';

const { Content } = Layout;

export interface IAppLayoutProps {
  routes: Array<RouteDescriptor>;
  TooltipButtonData?: TooltipButtonData;
  TooltipButtonLink?: string;
  transparentNavbar?: boolean;
}

const AppLayout: FC<IAppLayoutProps> = ({ ...props }) => {
  const [sideLeftShow, setSideLeftShow] = useState(false);
  const { routes, transparentNavbar, TooltipButtonData, TooltipButtonLink } =
    props;
  return (
    <BrowserRouter basename={PUBLIC_URL}>
      <Layout className="h-full">
        <Navbar
          logoutHandler={logout}
          routes={routes}
          transparent={transparentNavbar}
        />
        <Content className="flex">
          <Switch>
            {routes.map(r =>
              r.content ? (
                <Route exact key={r.route.path} path={r.route.path}>
                  <Row className="h-full pt-5 xs:pt-10 pb-20 flex w-full px-4">
                    <Col span={0} lg={1} xxl={2}></Col>
                    {r.content}
                    <Col span={0} lg={1} xxl={2}></Col>
                  </Row>
                </Route>
              ) : null
            )}
            <Route>
              <div className="flex justify-center items-center w-full">
                <Result
                  status="404"
                  title="404"
                  subTitle="Sorry, the page you visited does not exist."
                />
              </div>
            </Route>
          </Switch>
        </Content>
        <div className="right-TooltipButton flex">
          <TooltipButton
            TooltipButtonData={{
              tooltipTitle: 'Show CrownLabs infos',
              tooltipPlacement: 'right',
              type: 'primary',
              icon: <InfoOutlined style={{ fontSize: '22px' }} />,
            }}
            onClick={() => setSideLeftShow(true)}
          />
          {TooltipButtonData && (
            <TooltipButton
              className="ml-4"
              TooltipButtonData={{
                tooltipTitle: TooltipButtonData.tooltipTitle,
                tooltipPlacement: TooltipButtonData.tooltipPlacement,
                type: TooltipButtonData.type,
                icon: TooltipButtonData.icon,
              }}
              onClick={() => window.open(TooltipButtonLink, '_blank')}
            />
          )}
        </div>

        <div className="left-TooltipButton flex justify-between w-40">
          <Tooltip title="Virtual CPU usage" placement="topLeft">
            <Progress
              type="circle"
              format={p => <div>{`CPU\n${p}%`}</div>}
              percent={100}
              status="active"
              strokeColor={'#a61d24'}
              width={43}
              strokeWidth={10}
            />
          </Tooltip>
          <Tooltip title="Memory usage" placement="top">
            <Progress
              type="circle"
              percent={70}
              status="active"
              format={p => <div>{`RAM\n${p}%`}</div>}
              width={43}
              strokeWidth={10}
            />
          </Tooltip>
          <Tooltip title="Instances used" placement="top">
            <Progress
              type="circle"
              percent={75}
              width={43}
              status="active"
              format={() => <div className="align-middle text-center">5/5</div>}
              strokeWidth={10}
            />
          </Tooltip>
        </div>
      </Layout>
      <SidebarInfo
        show={sideLeftShow}
        setShow={setSideLeftShow}
        position="left"
      />
    </BrowserRouter>
  );
};

export default AppLayout;
