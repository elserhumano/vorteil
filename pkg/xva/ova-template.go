package xva

const ovaXMLTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<value>
	<struct>
		<member>
			<name>version</name>
			<value>
				<struct>
					<member>
						<name>hostname</name>
						<value>519e6dbb346d</value>
					</member>
					<member>
						<name>date</name>
						<value>2017-02-16</value>
					</member>
					<member>
						<name>product_version</name>
						<value>7.1.0</value>
					</member>
					<member>
						<name>product_brand</name>
						<value>XenServer</value>
					</member>
					<member>
						<name>build_number</name>
						<value>137272c</value>
					</member>
					<member>
						<name>xapi_major</name>
						<value>1</value>
					</member>
					<member>
						<name>xapi_minor</name>
						<value>9</value>
					</member>
					<member>
						<name>export_vsn</name>
						<value>2</value>
					</member>
				</struct>
			</value>
		</member>
		<member>
			<name>objects</name>
			<value>
				<array>
					<data>
						<value>
							<struct>
								<member>
									<name>class</name>
									<value>VM</value>
								</member>
								<member>
									<name>id</name>
									<value>Ref:0</value>
								</member>
								<member>
									<name>snapshot</name>
									<value>
										<struct>
											<member>
												<name>uuid</name>
												<value></value>
											</member>
											<member>
												<name>power_state</name>
												<value>Halted</value>
											</member>
											<member>
												<name>name_label</name>
												<value>%s</value>
											</member>
											<member>
												<name>name_description</name>
												<value>%s</value>
											</member>
											<member>
												<name>user_version</name>
												<value>1</value>
											</member>
											<member>
												<name>is_a_template</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>suspend_VDI</name>
												<value>OpaqueRef:NULL</value>
											</member>
											<member>
												<name>resident_on</name>
												<value>Ref:7</value>
											</member>
											<member>
												<name>affinity</name>
												<value>Ref:8</value>
											</member>
											<member>
												<name>memory_overhead</name>
												<value>0</value>
											</member>
											<member>
												<name>memory_target</name>
												<value>0</value>
											</member>
											<member>
												<name>memory_static_max</name>
												<value>%d</value>
											</member>
											<member>
												<name>memory_dynamic_max</name>
												<value>%d</value>
											</member>
											<member>
												<name>memory_dynamic_min</name>
												<value>%d</value>
											</member>
											<member>
												<name>memory_static_min</name>
												<value>%d</value>
											</member>
											<member>
												<name>VCPUs_params</name>
												<value>
													<struct />
												</value>
											</member>
											<member>
												<name>VCPUs_max</name>
												<value>%d</value>
											</member>
											<member>
												<name>VCPUs_at_startup</name>
												<value>%d</value>
											</member>
											<member>
												<name>actions_after_shutdown</name>
												<value>destroy</value>
											</member>
											<member>
												<name>actions_after_reboot</name>
												<value>restart</value>
											</member>
											<member>
												<name>actions_after_crash</name>
												<value>restart</value>
											</member>
											<member>
												<name>consoles</name>
												<value>
													<array>
														<data />
													</array>
												</value>
											</member>
											<member>
												<name>VIFs</name>
												<value>
													<array><data>%s</data></array>
												</value>
											</member>
											<member>
												<name>VBDs</name>
												<value>
													<array>
														<data>
															<value>Ref:3</value>
														</data>
													</array>
												</value>
											</member>
											<member>
												<name>VTPMs</name>
												<value>
													<array>
														<data />
													</array>
												</value>
											</member>
											<member>
												<name>PV_bootloader</name>
												<value />
											</member>
											<member>
												<name>PV_kernel</name>
												<value />
											</member>
											<member>
												<name>PV_ramdisk</name>
												<value />
											</member>
											<member>
												<name>PV_args</name>
												<value />
											</member>
											<member>
												<name>PV_bootloader_args</name>
												<value />
											</member>
											<member>
												<name>PV_legacy_args</name>
												<value />
											</member>
											<member>
												<name>HVM_boot_policy</name>
												<value>BIOS order</value>
											</member>
											<member>
												<name>HVM_boot_params</name>
												<value>
													<struct>
														<member>
															<name>order</name>
															<value>dc</value>
														</member>
													</struct>
												</value>
											</member>
											<member>
												<name>HVM_shadow_multiplier</name>
												<value>
													<double>1</double>
												</value>
											</member>
											<member>
												<name>platform</name>
												<value>
													<struct>
														<member>
															<name>timeoffset</name>
															<value>0</value>
														</member>
														<member>
															<name>stdvga</name>
															<value>0</value>
														</member>
														<member>
															<name>apic</name>
															<value>true</value>
														</member>
														<member>
															<name>acpi</name>
															<value>true</value>
														</member>
														<member>
															<name>nx</name>
															<value>true</value>
														</member>
														<member>
															<name>pae</name>
															<value>true</value>
														</member>
													</struct>
												</value>
											</member>
											<member>
												<name>PCI_bus</name>
												<value />
											</member>
											<member>
												<name>other_config</name>
												<value>
													<struct>
														<member>
															<name>vgpu_pci</name>
															<value />
														</member>
														<member>
															<name>mac_seed</name>
															<value>890231c4-c804-44f3-efa5-fa6ec0719286</value>
														</member>
													</struct>
												</value>
											</member>
											<member>
												<name>domid</name>
												<value>-1</value>
											</member>
											<member>
												<name>domarch</name>
												<value />
											</member>
											<member>
												<name>last_boot_CPU_flags</name>
												<value>
													<struct>
														<member>
															<name>vendor</name>
															<value>GenuineIntel</value>
														</member>
														<member>
															<name>features</name>
															<value>178bfbff-f7fa3223-2d93fbff-00000023-00000001-000007ab-00000000-00000000-00000000</value>
														</member>
													</struct>
												</value>
											</member>
											<member>
												<name>is_control_domain</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>metrics</name>
												<value>OpaqueRef:NULL</value>
											</member>
											<member>
												<name>guest_metrics</name>
												<value>Ref:6</value>
											</member>
											<member>
												<name>last_booted_record</name>
												<value>('struct' ('uuid' 'd42c38db-76ca-8e53-a0c1-1f0cae4d0ad9') ('allowed_operations' ('array')) ('current_operations' ('struct' ('OpaqueRef:a664b26b-3b87-83e7-90ca-9a784caccc02' 'start'))) ('power_state' 'Halted') ('name_label' 'ddd') ('name_description' 'Created by XenCenter Disk Image Import') ('user_version' '1') ('is_a_template' ('boolean' '0')) ('suspend_VDI' 'OpaqueRef:NULL') ('resident_on' 'OpaqueRef:NULL') ('affinity' '') ('memory_overhead' '7340032') ('memory_target' '0') ('memory_static_max' '536870912') ('memory_dynamic_max' '536870912') ('memory_dynamic_min' '536870912') ('memory_static_min' '16777216') ('VCPUs_params' ('struct')) ('VCPUs_max' '1') ('VCPUs_at_startup' '1') ('actions_after_shutdown' 'destroy') ('actions_after_reboot' 'restart') ('actions_after_crash' 'restart') ('consoles' ('array')) ('VIFs' ('array' 'OpaqueRef:17362699-5416-fdf8-2ebc-1469480a3697')) ('VBDs' ('array' 'OpaqueRef:4ae40d83-6373-8ea2-2801-2385caba0df2')) ('crash_dumps' ('array')) ('VTPMs' ('array')) ('PV_bootloader' '') ('PV_kernel' '') ('PV_ramdisk' '') ('PV_args' '') ('PV_bootloader_args' '') ('PV_legacy_args' '') ('HVM_boot_policy' 'BIOS order') ('HVM_boot_params' ('struct' ('order' 'dc'))) ('HVM_shadow_multiplier' ('double' '1')) ('platform' ('struct' ('timeoffset' '0') ('stdvga' '0') ('apic' 'true') ('acpi' 'true') ('nx' 'true') ('pae' 'true'))) ('PCI_bus' '') ('other_config' ('struct' ('vgpu_pci' '') ('mac_seed' '890231c4-c804-44f3-efa5-fa6ec0719286'))) ('domid' '-1') ('domarch' '') ('last_boot_CPU_flags' ('struct' ('vendor' 'GenuineIntel') ('features' '178bfbff-f7fa3223-2d93fbff-00000023-00000001-000007ab-00000000-00000000-00000000'))) ('is_control_domain' ('boolean' '0')) ('metrics' 'OpaqueRef:a1055d02-9de5-acf4-c860-c560c5f04ec0') ('guest_metrics' 'OpaqueRef:460f13a2-2591-5671-3bc7-f09a5e9a1c83') ('last_booted_record' '') ('recommendations' '') ('xenstore_data' ('struct' ('vm-data' ''))) ('ha_always_run' ('boolean' '0')) ('ha_restart_priority' '') ('is_a_snapshot' ('boolean' '0')) ('snapshot_of' 'OpaqueRef:NULL') ('snapshots' ('array')) ('snapshot_time' ('dateTime.iso8601' '19700101T00:00:00Z')) ('transportable_snapshot_id' '') ('blobs' ('struct')) ('tags' ('array')) ('blocked_operations' ('struct')) ('snapshot_info' ('struct')) ('snapshot_metadata' '') ('parent' 'OpaqueRef:NULL') ('children' ('array')) ('bios_strings' ('struct')) ('protection_policy' 'OpaqueRef:NULL') ('is_snapshot_from_vmpp' ('boolean' '0')) ('appliance' '') ('start_delay' '0') ('shutdown_delay' '0') ('order' '0') ('VGPUs' ('array')) ('attached_PCIs' ('array')) ('suspend_SR' '') ('version' '0') ('generation_id' '') ('hardware_platform_version' '0') ('has_vendor_device' ('boolean' '0')) ('requires_reboot' ('boolean' '0')) ('reference_label' ''))</value>
											</member>
											<member>
												<name>recommendations</name>
												<value />
											</member>
											<member>
												<name>xenstore_data</name>
												<value>
													<struct>
														<member>
															<name>vm-data</name>
															<value />
														</member>
													</struct>
												</value>
											</member>
											<member>
												<name>ha_always_run</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>ha_restart_priority</name>
												<value />
											</member>
											<member>
												<name>is_a_snapshot</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>snapshot_of</name>
												<value>OpaqueRef:NULL</value>
											</member>
											<member>
												<name>snapshots</name>
												<value>
													<array>
														<data />
													</array>
												</value>
											</member>
											<member>
												<name>snapshot_time</name>
												<value>
													<dateTime.iso8601>19700101T00:00:00Z</dateTime.iso8601>
												</value>
											</member>
											<member>
												<name>transportable_snapshot_id</name>
												<value />
											</member>
											<member>
												<name>blobs</name>
												<value>
													<struct />
												</value>
											</member>
											<member>
												<name>blocked_operations</name>
												<value>
													<struct />
												</value>
											</member>
											<member>
												<name>snapshot_info</name>
												<value>
													<struct />
												</value>
											</member>
											<member>
												<name>snapshot_metadata</name>
												<value />
											</member>
											<member>
												<name>parent</name>
												<value>OpaqueRef:NULL</value>
											</member>
											<member>
												<name>children</name>
												<value>
													<array>
														<data />
													</array>
												</value>
											</member>
											<member>
												<name>bios_strings</name>
												<value>
													<struct>
														<member>
															<name>bios-vendor</name>
															<value>Xen</value>
														</member>
														<member>
															<name>bios-version</name>
															<value />
														</member>
														<member>
															<name>system-manufacturer</name>
															<value>Xen</value>
														</member>
														<member>
															<name>system-product-name</name>
															<value>HVM domU</value>
														</member>
														<member>
															<name>system-version</name>
															<value />
														</member>
														<member>
															<name>system-serial-number</name>
															<value />
														</member>
														<member>
															<name>hp-rombios</name>
															<value />
														</member>
														<member>
															<name>oem-1</name>
															<value>Xen</value>
														</member>
														<member>
															<name>oem-2</name>
															<value>MS_VM_CERT/SHA1/bdbeb6e0a816d43fa6d3fe8aaef04c2bad9d3e3d</value>
														</member>
													</struct>
												</value>
											</member>
											<member>
												<name>protection_policy</name>
												<value>OpaqueRef:NULL</value>
											</member>
											<member>
												<name>is_snapshot_from_vmpp</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>appliance</name>
												<value />
											</member>
											<member>
												<name>start_delay</name>
												<value>0</value>
											</member>
											<member>
												<name>shutdown_delay</name>
												<value>0</value>
											</member>
											<member>
												<name>order</name>
												<value>0</value>
											</member>
											<member>
												<name>VGPUs</name>
												<value>
													<array>
														<data />
													</array>
												</value>
											</member>
											<member>
												<name>attached_PCIs</name>
												<value>
													<array>
														<data />
													</array>
												</value>
											</member>
											<member>
												<name>suspend_SR</name>
												<value />
											</member>
											<member>
												<name>version</name>
												<value>0</value>
											</member>
											<member>
												<name>generation_id</name>
												<value />
											</member>
											<member>
												<name>hardware_platform_version</name>
												<value>0</value>
											</member>
											<member>
												<name>has_vendor_device</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>requires_reboot</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>reference_label</name>
												<value />
											</member>
										</struct>
									</value>
								</member>
							</struct>
						</value>
						<value>
							<struct>
								<member>
									<name>class</name>
									<value>VBD</value>
								</member>
								<member>
									<name>id</name>
									<value>Ref:3</value>
								</member>
								<member>
									<name>snapshot</name>
									<value>
										<struct>
											<member>
												<name>uuid</name>
												<value></value>
											</member>
											<member>
												<name>VM</name>
												<value>Ref:0</value>
											</member>
											<member>
												<name>VDI</name>
												<value>Ref:4</value>
											</member>
											<member>
												<name>device</name>
												<value>xvda</value>
											</member>
											<member>
												<name>userdevice</name>
												<value>0</value>
											</member>
											<member>
												<name>bootable</name>
												<value>
													<boolean>1</boolean>
												</value>
											</member>
											<member>
												<name>mode</name>
												<value>RW</value>
											</member>
											<member>
												<name>type</name>
												<value>Disk</value>
											</member>
											<member>
												<name>unpluggable</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>storage_lock</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>empty</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>other_config</name>
												<value>
													<struct />
												</value>
											</member>
											<member>
												<name>currently_attached</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>status_code</name>
												<value>0</value>
											</member>
											<member>
												<name>status_detail</name>
												<value />
											</member>
											<member>
												<name>runtime_properties</name>
												<value>
													<struct />
												</value>
											</member>
											<member>
												<name>qos_algorithm_type</name>
												<value />
											</member>
											<member>
												<name>qos_algorithm_params</name>
												<value>
													<struct />
												</value>
											</member>
											<member>
												<name>qos_supported_algorithms</name>
												<value>
													<array>
														<data />
													</array>
												</value>
											</member>
											<member>
												<name>metrics</name>
												<value>OpaqueRef:NULL</value>
											</member>
										</struct>
									</value>
								</member>
							</struct>
						</value>%s
						<value>
							<struct>
								<member>
									<name>class</name>
									<value>VDI</value>
								</member>
								<member>
									<name>id</name>
									<value>Ref:4</value>
								</member>
								<member>
									<name>snapshot</name>
									<value>
										<struct>
											<member>
												<name>uuid</name>
												<value></value>
											</member>
											<member>
												<name>SR</name>
												<value>Ref:5</value>
											</member>
											<member>
												<name>virtual_size</name>
												<value>%d</value>
											</member>
											<member>
												<name>physical_utilisation</name>
												<value>0</value>
											</member>
											<member>
												<name>type</name>
												<value>user</value>
											</member>
											<member>
												<name>sharable</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>read_only</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>other_config</name>
												<value>
													<struct />
												</value>
											</member>
											<member>
												<name>storage_lock</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>managed</name>
												<value>
													<boolean>1</boolean>
												</value>
											</member>
											<member>
												<name>missing</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>parent</name>
												<value>OpaqueRef:NULL</value>
											</member>
										</struct>
									</value>
								</member>
							</struct>
						</value>
						<value>
							<struct>
								<member>
									<name>class</name>
									<value>SR</value>
								</member>
								<member>
									<name>id</name>
									<value>Ref:5</value>
								</member>
								<member>
									<name>snapshot</name>
									<value>
										<struct>
											<member>
												<name>uuid</name>
												<value></value>
											</member>
											<member>
												<name>virtual_allocation</name>
												<value>0</value>
											</member>
											<member>
												<name>physical_utilisation</name>
												<value>0</value>
											</member>
											<member>
												<name>physical_size</name>
												<value>0</value>
											</member>
											<member>
												<name>type</name>
												<value>lvm</value>
											</member>
											<member>
												<name>content_type</name>
												<value>user</value>
											</member>
											<member>
												<name>shared</name>
												<value>
													<boolean>0</boolean>
												</value>
											</member>
											<member>
												<name>other_config</name>
												<value>
													<struct />
												</value>
											</member>
										</struct>
									</value>
								</member>
							</struct>
						</value>
					</data>
				</array>
			</value>
		</member>
	</struct>
</value>`
