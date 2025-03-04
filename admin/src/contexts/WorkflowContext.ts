import React from "react";
import { GetWorkflowResponse } from "../api/Workflow";
import { WorkflowTabType } from "../models/WorkflowTab";

export const WorkflowContext = React.createContext({
  refresh: () => {},
  data: null as GetWorkflowResponse | null,
  name: "",
  tab: WorkflowTabType.Status,
  group: "",
});
