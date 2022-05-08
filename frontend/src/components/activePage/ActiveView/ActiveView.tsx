import { DeleteOutlined } from '@ant-design/icons';
import { Col, Progress, Space, Tooltip } from 'antd';
import Button from 'antd-button-color';
import { FC, useEffect, useState } from 'react';
import { User, WorkspaceRole } from '../../../utils';
import { SessionValue, StorageKeys } from '../../../utilsStorage';
import Box from '../../common/Box';
import { ModalAlert } from '../../common/ModalAlert';
import TableInstanceLogic from '../TableInstance/TableInstanceLogic';
import TableWorkspaceLogic from '../TableWorkspaceLogic/TableWorkspaceLogic';
import Toolbox from '../Toolbox/Toolbox';
import ViewModeButton from './ViewModeButton/ViewModeButton';

const view = new SessionValue(StorageKeys.Active_View, WorkspaceRole.user);
const advanced = new SessionValue(StorageKeys.Active_Headers, 'true');

export interface IActiveViewProps {
  user: User;
  workspaces: Array<{
    prettyName: string;
    role: WorkspaceRole;
    namespace: string;
    id: string;
  }>;
  managerView: boolean;
}

const ActiveView: FC<IActiveViewProps> = ({ ...props }) => {
  const { managerView, user, workspaces } = props;
  const [expandAll, setExpandAll] = useState(false);
  const [collapseAll, setCollapseAll] = useState(false);
  const [destroySelectedTrigger, setDestroySelectedTrigger] = useState(false);
  const [showAlert, setShowAlert] = useState(false);
  const [searchField, setSearchField] = useState('');
  const [currentView, setCurrentView] = useState<WorkspaceRole>(
    managerView ? (view.get() as WorkspaceRole) : WorkspaceRole.user
  );
  const [showAdvanced, setShowAdvanced] = useState(
    !managerView || advanced.get() !== 'false'
  );
  const [showCheckbox, setShowCheckbox] = useState(false);
  const [selectiveDestroy, setSelectiveDestroy] = useState<string[]>([]);

  const selectToDestroy = (instanceId: string) => {
    selectiveDestroy.includes(instanceId)
      ? setSelectiveDestroy(old => old.filter(id => id !== instanceId))
      : setSelectiveDestroy(old => [...old, instanceId]);
  };

  const deselectAll = () => setSelectiveDestroy([]);

  const displayCheckbox = () => {
    if (!showCheckbox) {
      setShowCheckbox(true);
    } else {
      setShowCheckbox(() => {
        deselectAll();
        return false;
      });
    }
  };

  useEffect(() => {
    view.set(currentView);
  }, [currentView]);

  useEffect(() => {
    advanced.set(String(showAdvanced));
  }, [showAdvanced]);

  return (
    <Col span={24} lg={22} xxl={20}>
      <ModalAlert
        headTitle="Destroy Selected"
        show={showAlert}
        message="ATTENTION"
        description={`Are you sure do you want to destroy the ${selectiveDestroy.length} selected instances. This operation is dangerous and irreversible!`}
        type="error"
        buttons={[
          <Button
            type="danger"
            shape="round"
            size="middle"
            icon={<DeleteOutlined />}
            className="border-0"
            onClick={() => {
              setDestroySelectedTrigger(true);
              setShowAlert(false);
            }}
          >
            Destroy Selected
          </Button>,
        ]}
        setShow={setShowAlert}
      />
      <Box
        header={{
          size: 'middle',
          right: managerView && (
            <div className="h-full flex justify-center items-center pr-10">
              <Space size="small">
                <ViewModeButton
                  setCurrentView={setCurrentView}
                  currentView={currentView}
                />
              </Space>
            </div>
          ),
          left: (
            <>
              {managerView && currentView === WorkspaceRole.manager ? (
                <div className="h-full flex justify-center items-center pl-6 gap-4">
                  <Toolbox
                    setSearchField={setSearchField}
                    setExpandAll={setExpandAll}
                    setCollapseAll={setCollapseAll}
                    showAdvanced={showAdvanced}
                    setShowAdvanced={setShowAdvanced}
                    showCheckbox={showCheckbox}
                    setShowCheckbox={displayCheckbox}
                    setShowAlert={setShowAlert}
                    selectiveDestroy={selectiveDestroy}
                    deselectAll={deselectAll}
                  />
                </div>
              ) : (
                <div className="h-full flex justify-center items-center pl-6 gap-4">
                  <Tooltip title="Virtual CPU usage" placement="bottom">
                    <Progress
                      type="circle"
                      format={p => <div>{`CPU\n${p}%`}</div>}
                      percent={100}
                      status="active"
                      strokeColor={'#a61d24'}
                      width={50}
                      strokeWidth={10}
                    />
                  </Tooltip>
                  <Tooltip title="Memory usage" placement="bottom">
                    <Progress
                      type="circle"
                      percent={70}
                      status="active"
                      format={p => <div>{`RAM\n${p}%`}</div>}
                      width={50}
                      strokeWidth={10}
                    />
                  </Tooltip>
                  <Tooltip title="Instances used" placement="bottom">
                    <Progress
                      type="circle"
                      percent={(4 / 5) * 100}
                      width={50}
                      status="active"
                      format={p => (
                        <div className="align-middle text-center">
                          {(p! / 100) * 5}/5
                        </div>
                      )}
                      strokeWidth={10}
                    />
                  </Tooltip>
                </div>
              )}
            </>
          ),
        }}
      >
        {currentView === WorkspaceRole.manager && managerView ? (
          <div className="flex flex-col justify-start">
            <TableWorkspaceLogic
              workspaces={workspaces}
              user={user}
              filter={searchField}
              collapseAll={collapseAll}
              expandAll={expandAll}
              setCollapseAll={setCollapseAll}
              setExpandAll={setExpandAll}
              showAdvanced={showAdvanced}
              showCheckbox={showCheckbox}
              destroySelectedTrigger={destroySelectedTrigger}
              setDestroySelectedTrigger={setDestroySelectedTrigger}
              selectiveDestroy={selectiveDestroy}
              selectToDestroy={selectToDestroy}
            />
          </div>
        ) : (
          <TableInstanceLogic
            showGuiIcon={true}
            user={user}
            viewMode={currentView}
            extended={true}
          />
        )}
      </Box>
    </Col>
  );
};

export default ActiveView;
