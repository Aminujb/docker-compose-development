import { Form, Input, Button } from "antd";
import axios from "axios";
import { useHistory } from "react-router-dom";

const layout = {
  labelCol: {
    span: 8,
  },
  wrapperCol: {
    span: 16,
  },
};
/* eslint-disable no-template-curly-in-string */

const validateMessages = {
  required: "${label} is required!",
  types: {
    email: "${label} is not a valid email!",
    number: "${label} is not a valid number!",
  },
  number: {
    range: "${label} must be between ${min} and ${max}",
  },
};
/* eslint-enable no-template-curly-in-string */

const CreateReport = () => {
  const history = useHistory();

  const onFinish = (values) => {
    const requestBody = {
      username: values.report.username,
      report: {
        yesterday: values.report.yesterday,
        today: values.report.today,
        blockers: values.report.blockers,
      },
    };

    axios
      .post(`${process.env.REACT_APP_BACKEND_URL}/daily-standup`, requestBody)
      .then((res) => {
        history.push("/");
      });
  };

  return (
    <div style={{ marginRight: "80px" }}>
      <Form
        {...layout}
        name="nest-messages"
        onFinish={onFinish}
        validateMessages={validateMessages}
      >
        <Form.Item
          name={["report", "username"]}
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
          name={["report", "yesterday"]}
          label="What did you do yesterday?"
          rules={[
            {
              required: true,
            },
          ]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          name={["report", "today"]}
          label="What will you do today?"
          rules={[
            {
              required: true,
            },
          ]}
        >
          <Input />
        </Form.Item>
        <Form.Item
          name={["report", "blockers"]}
          label="Anything blocking your progress?"
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
    </div>
  );
};

export default CreateReport;
