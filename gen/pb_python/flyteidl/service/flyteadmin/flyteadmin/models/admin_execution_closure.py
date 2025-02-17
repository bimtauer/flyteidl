# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from flyteadmin.models.admin_abort_metadata import AdminAbortMetadata  # noqa: F401,E501
from flyteadmin.models.admin_literal_map_blob import AdminLiteralMapBlob  # noqa: F401,E501
from flyteadmin.models.admin_notification import AdminNotification  # noqa: F401,E501
from flyteadmin.models.core_execution_error import CoreExecutionError  # noqa: F401,E501
from flyteadmin.models.core_identifier import CoreIdentifier  # noqa: F401,E501
from flyteadmin.models.core_literal_map import CoreLiteralMap  # noqa: F401,E501
from flyteadmin.models.core_workflow_execution_phase import CoreWorkflowExecutionPhase  # noqa: F401,E501


class AdminExecutionClosure(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'outputs': 'AdminLiteralMapBlob',
        'error': 'CoreExecutionError',
        'abort_cause': 'str',
        'abort_metadata': 'AdminAbortMetadata',
        'output_data': 'CoreLiteralMap',
        'computed_inputs': 'CoreLiteralMap',
        'phase': 'CoreWorkflowExecutionPhase',
        'started_at': 'datetime',
        'duration': 'str',
        'created_at': 'datetime',
        'updated_at': 'datetime',
        'notifications': 'list[AdminNotification]',
        'workflow_id': 'CoreIdentifier'
    }

    attribute_map = {
        'outputs': 'outputs',
        'error': 'error',
        'abort_cause': 'abort_cause',
        'abort_metadata': 'abort_metadata',
        'output_data': 'output_data',
        'computed_inputs': 'computed_inputs',
        'phase': 'phase',
        'started_at': 'started_at',
        'duration': 'duration',
        'created_at': 'created_at',
        'updated_at': 'updated_at',
        'notifications': 'notifications',
        'workflow_id': 'workflow_id'
    }

    def __init__(self, outputs=None, error=None, abort_cause=None, abort_metadata=None, output_data=None, computed_inputs=None, phase=None, started_at=None, duration=None, created_at=None, updated_at=None, notifications=None, workflow_id=None):  # noqa: E501
        """AdminExecutionClosure - a model defined in Swagger"""  # noqa: E501

        self._outputs = None
        self._error = None
        self._abort_cause = None
        self._abort_metadata = None
        self._output_data = None
        self._computed_inputs = None
        self._phase = None
        self._started_at = None
        self._duration = None
        self._created_at = None
        self._updated_at = None
        self._notifications = None
        self._workflow_id = None
        self.discriminator = None

        if outputs is not None:
            self.outputs = outputs
        if error is not None:
            self.error = error
        if abort_cause is not None:
            self.abort_cause = abort_cause
        if abort_metadata is not None:
            self.abort_metadata = abort_metadata
        if output_data is not None:
            self.output_data = output_data
        if computed_inputs is not None:
            self.computed_inputs = computed_inputs
        if phase is not None:
            self.phase = phase
        if started_at is not None:
            self.started_at = started_at
        if duration is not None:
            self.duration = duration
        if created_at is not None:
            self.created_at = created_at
        if updated_at is not None:
            self.updated_at = updated_at
        if notifications is not None:
            self.notifications = notifications
        if workflow_id is not None:
            self.workflow_id = workflow_id

    @property
    def outputs(self):
        """Gets the outputs of this AdminExecutionClosure.  # noqa: E501

        Output URI in the case of a successful execution.  # noqa: E501

        :return: The outputs of this AdminExecutionClosure.  # noqa: E501
        :rtype: AdminLiteralMapBlob
        """
        return self._outputs

    @outputs.setter
    def outputs(self, outputs):
        """Sets the outputs of this AdminExecutionClosure.

        Output URI in the case of a successful execution.  # noqa: E501

        :param outputs: The outputs of this AdminExecutionClosure.  # noqa: E501
        :type: AdminLiteralMapBlob
        """

        self._outputs = outputs

    @property
    def error(self):
        """Gets the error of this AdminExecutionClosure.  # noqa: E501

        Error information in the case of a failed execution.  # noqa: E501

        :return: The error of this AdminExecutionClosure.  # noqa: E501
        :rtype: CoreExecutionError
        """
        return self._error

    @error.setter
    def error(self, error):
        """Sets the error of this AdminExecutionClosure.

        Error information in the case of a failed execution.  # noqa: E501

        :param error: The error of this AdminExecutionClosure.  # noqa: E501
        :type: CoreExecutionError
        """

        self._error = error

    @property
    def abort_cause(self):
        """Gets the abort_cause of this AdminExecutionClosure.  # noqa: E501

        In the case of a user-specified abort, this will pass along the user-supplied cause.  # noqa: E501

        :return: The abort_cause of this AdminExecutionClosure.  # noqa: E501
        :rtype: str
        """
        return self._abort_cause

    @abort_cause.setter
    def abort_cause(self, abort_cause):
        """Sets the abort_cause of this AdminExecutionClosure.

        In the case of a user-specified abort, this will pass along the user-supplied cause.  # noqa: E501

        :param abort_cause: The abort_cause of this AdminExecutionClosure.  # noqa: E501
        :type: str
        """

        self._abort_cause = abort_cause

    @property
    def abort_metadata(self):
        """Gets the abort_metadata of this AdminExecutionClosure.  # noqa: E501

        In the case of a user-specified abort, this will pass along the user and their supplied cause.  # noqa: E501

        :return: The abort_metadata of this AdminExecutionClosure.  # noqa: E501
        :rtype: AdminAbortMetadata
        """
        return self._abort_metadata

    @abort_metadata.setter
    def abort_metadata(self, abort_metadata):
        """Sets the abort_metadata of this AdminExecutionClosure.

        In the case of a user-specified abort, this will pass along the user and their supplied cause.  # noqa: E501

        :param abort_metadata: The abort_metadata of this AdminExecutionClosure.  # noqa: E501
        :type: AdminAbortMetadata
        """

        self._abort_metadata = abort_metadata

    @property
    def output_data(self):
        """Gets the output_data of this AdminExecutionClosure.  # noqa: E501

        Raw output data produced by this execution.  # noqa: E501

        :return: The output_data of this AdminExecutionClosure.  # noqa: E501
        :rtype: CoreLiteralMap
        """
        return self._output_data

    @output_data.setter
    def output_data(self, output_data):
        """Sets the output_data of this AdminExecutionClosure.

        Raw output data produced by this execution.  # noqa: E501

        :param output_data: The output_data of this AdminExecutionClosure.  # noqa: E501
        :type: CoreLiteralMap
        """

        self._output_data = output_data

    @property
    def computed_inputs(self):
        """Gets the computed_inputs of this AdminExecutionClosure.  # noqa: E501


        :return: The computed_inputs of this AdminExecutionClosure.  # noqa: E501
        :rtype: CoreLiteralMap
        """
        return self._computed_inputs

    @computed_inputs.setter
    def computed_inputs(self, computed_inputs):
        """Sets the computed_inputs of this AdminExecutionClosure.


        :param computed_inputs: The computed_inputs of this AdminExecutionClosure.  # noqa: E501
        :type: CoreLiteralMap
        """

        self._computed_inputs = computed_inputs

    @property
    def phase(self):
        """Gets the phase of this AdminExecutionClosure.  # noqa: E501

        Most recent recorded phase for the execution.  # noqa: E501

        :return: The phase of this AdminExecutionClosure.  # noqa: E501
        :rtype: CoreWorkflowExecutionPhase
        """
        return self._phase

    @phase.setter
    def phase(self, phase):
        """Sets the phase of this AdminExecutionClosure.

        Most recent recorded phase for the execution.  # noqa: E501

        :param phase: The phase of this AdminExecutionClosure.  # noqa: E501
        :type: CoreWorkflowExecutionPhase
        """

        self._phase = phase

    @property
    def started_at(self):
        """Gets the started_at of this AdminExecutionClosure.  # noqa: E501

        Reported time at which the execution began running.  # noqa: E501

        :return: The started_at of this AdminExecutionClosure.  # noqa: E501
        :rtype: datetime
        """
        return self._started_at

    @started_at.setter
    def started_at(self, started_at):
        """Sets the started_at of this AdminExecutionClosure.

        Reported time at which the execution began running.  # noqa: E501

        :param started_at: The started_at of this AdminExecutionClosure.  # noqa: E501
        :type: datetime
        """

        self._started_at = started_at

    @property
    def duration(self):
        """Gets the duration of this AdminExecutionClosure.  # noqa: E501

        The amount of time the execution spent running.  # noqa: E501

        :return: The duration of this AdminExecutionClosure.  # noqa: E501
        :rtype: str
        """
        return self._duration

    @duration.setter
    def duration(self, duration):
        """Sets the duration of this AdminExecutionClosure.

        The amount of time the execution spent running.  # noqa: E501

        :param duration: The duration of this AdminExecutionClosure.  # noqa: E501
        :type: str
        """

        self._duration = duration

    @property
    def created_at(self):
        """Gets the created_at of this AdminExecutionClosure.  # noqa: E501

        Reported time at which the execution was created.  # noqa: E501

        :return: The created_at of this AdminExecutionClosure.  # noqa: E501
        :rtype: datetime
        """
        return self._created_at

    @created_at.setter
    def created_at(self, created_at):
        """Sets the created_at of this AdminExecutionClosure.

        Reported time at which the execution was created.  # noqa: E501

        :param created_at: The created_at of this AdminExecutionClosure.  # noqa: E501
        :type: datetime
        """

        self._created_at = created_at

    @property
    def updated_at(self):
        """Gets the updated_at of this AdminExecutionClosure.  # noqa: E501

        Reported time at which the execution was last updated.  # noqa: E501

        :return: The updated_at of this AdminExecutionClosure.  # noqa: E501
        :rtype: datetime
        """
        return self._updated_at

    @updated_at.setter
    def updated_at(self, updated_at):
        """Sets the updated_at of this AdminExecutionClosure.

        Reported time at which the execution was last updated.  # noqa: E501

        :param updated_at: The updated_at of this AdminExecutionClosure.  # noqa: E501
        :type: datetime
        """

        self._updated_at = updated_at

    @property
    def notifications(self):
        """Gets the notifications of this AdminExecutionClosure.  # noqa: E501

        The notification settings to use after merging the CreateExecutionRequest and the launch plan notification settings. An execution launched with notifications will always prefer that definition to notifications defined statically in a launch plan.  # noqa: E501

        :return: The notifications of this AdminExecutionClosure.  # noqa: E501
        :rtype: list[AdminNotification]
        """
        return self._notifications

    @notifications.setter
    def notifications(self, notifications):
        """Sets the notifications of this AdminExecutionClosure.

        The notification settings to use after merging the CreateExecutionRequest and the launch plan notification settings. An execution launched with notifications will always prefer that definition to notifications defined statically in a launch plan.  # noqa: E501

        :param notifications: The notifications of this AdminExecutionClosure.  # noqa: E501
        :type: list[AdminNotification]
        """

        self._notifications = notifications

    @property
    def workflow_id(self):
        """Gets the workflow_id of this AdminExecutionClosure.  # noqa: E501

        Identifies the workflow definition for this execution.  # noqa: E501

        :return: The workflow_id of this AdminExecutionClosure.  # noqa: E501
        :rtype: CoreIdentifier
        """
        return self._workflow_id

    @workflow_id.setter
    def workflow_id(self, workflow_id):
        """Sets the workflow_id of this AdminExecutionClosure.

        Identifies the workflow definition for this execution.  # noqa: E501

        :param workflow_id: The workflow_id of this AdminExecutionClosure.  # noqa: E501
        :type: CoreIdentifier
        """

        self._workflow_id = workflow_id

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(AdminExecutionClosure, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, AdminExecutionClosure):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
