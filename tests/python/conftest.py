"""Pulumi mock runtime for SDK smoke tests.

Set up before any pulumi_astronomer import so resource registrations
go through MockRuntime instead of trying to talk to a real engine.
"""

import pulumi


class _AstroMocks(pulumi.runtime.Mocks):
    def new_resource(self, args: pulumi.runtime.MockResourceArgs):
        # Echo inputs back as outputs; assign a deterministic ID derived from name.
        return [args.name + "_id", args.inputs]

    def call(self, args: pulumi.runtime.MockCallArgs):
        return {}


pulumi.runtime.set_mocks(_AstroMocks(), preview=False)
