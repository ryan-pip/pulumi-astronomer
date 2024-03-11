# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

organizationId: Optional[str]
"""
Organization id this provider will operate on.
"""

token: Optional[str]
"""
Astronomer API Token. Can be set with an `ASTRONOMER_API_TOKEN` env var.
"""

