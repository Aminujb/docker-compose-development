import React, { useState, useEffect } from "react";
import { Table, Space, Modal, Form, Input, Button } from "antd";
import axios from "axios";

const layout = {
  labelCol: {
    span: 8,
  },
  wrapperCol: {
    span: 14,
  },
};
/* eslint-disable no-template-curly-in-string */

const validateMessages = {
  required: "${label} is required!",
  types: {
    email: "${label} is not a valid email!",
    number: "${label} is not a valid number!",
  },
};
/* eslint-enable no-template-curly-in-string */

const Reports = () => {
  const [form] = Form.useForm();
  const [isModalVisible, setIsModalVisible] = useState(false);
  const [modalData, setModalData] = useState(null);

  function handleDeleteClick(id) {
    axios
      .delete(`${process.env.REACT_APP_BACKEND_URL}/daily-standup/${id}`)
      .then((res) => {
        window.location.reload(false);
      });
  }

  const onFinish = (values) => {
    const requestBody = {
      username: values.username,
      report: {
        yesterday: values.yesterday,
        today: values.today,
        blockers: values.blockers,
      },
    };

    axios
      .put(
        `${process.env.REACT_APP_BACKEND_URL}/daily-standup/${modalData.key}`,
        requestBody
      )
      .then((res) => {
        window.location.reload(false);
      });
  };

  const showModal = (data) => {
    setModalData(data);
    form.setFieldsValue({
      username: data.name,
      blockers: data.blockers,
      today: data.today,
      yesterday: data.yesterday,
    });
    setIsModalVisible(true);
  };

  const handleOk = () => {
    setIsModalVisible(false);
  };

  const handleCancel = () => {
    setIsModalVisible(false);
  };

  const columns = [
    {
      title: "Date",
      dataIndex: "date",
      key: "date",
      render: (text) => <a>{text}</a>,
    },
    {
      title: "Name",
      dataIndex: "name",
      key: "name",
      render: (text) => <a>{text}</a>,
    },
    {
      title: "What I did Yesterday",
      dataIndex: "yesterday",
      key: "yesterday",
    },
    {
      title: "What I Will do Today",
      dataIndex: "today",
      key: "today",
    },
    {
      title: "Blockers",
      key: "blockers",
      dataIndex: "blockers",
    },
    {
      title: "Action",
      key: "action",
      render: (text, record) => (
        <Space size="middle">
          <a onClick={() => showModal(record)}>Edit</a>
          <a onClick={() => handleDeleteClick(record.key)}>Delete</a>
        </Space>
      ),
    },
  ];

  const [reports, setReports] = useState([]);

  useEffect(() => {
    axios.get(`${process.env.REACT_APP_BACKEND_URL}/daily-standup`).then((res) => {
      setReports(res.data.data);
    });
  }, []);

  let data = null;

  if (reports !== null) {
    data = reports.map(function (report) {
      return {
        key: report._id,
        date: report.date_created,
        name: report.username,
        yesterday: report.report.yesterday,
        today: report.report.today,
        blockers: report.report.blockers,
      };
    });
  }

  return (
    <>
      <Table columns={columns} dataSource={data} />
      <Modal
        title="Edit Standup"
        visible={isModalVisible}
        onOk={handleOk}
        onCancel={handleCancel}
        okButtonProps={{ disabled: true }}
      >
        <Form
          {...layout}
          name="nest-messages"
          onFinish={onFinish}
          validateMessages={validateMessages}
          form={form}
        >
          <Form.Item
            name={["username"]}
            label="Username"
            rules={[
              {
                required: true,
              },
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name={["yesterday"]}
            label="Yesterday"
            rules={[
              {
                required: true,
              },
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name={["today"]}
            label="Today"
            rules={[
              {
                required: true,
              },
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name={["blockers"]}
            label="Blockers"
            rules={[
              {
                required: true,
              },
            ]}
          >
            <Input />
          </Form.Item>

          <Form.Item wrapperCol={{ ...layout.wrapperCol, offset: 8 }}>
            <Button type="primary" htmlType="submit">
              Submit
            </Button>
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default Reports;
